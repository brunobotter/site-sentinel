const API_BASE = "/monitor";

const table = document.getElementById("status-table");
const monitorForm = document.getElementById("monitor-form");
const feedback = document.getElementById("monitor-feedback");
const toggleFormBtn = document.getElementById("toggle-monitor-form");
const runChecksBtn = document.getElementById("run-checks");

const avgLatencyEl = document.getElementById("avg-latency");
const maxLatencyEl = document.getElementById("max-latency");
const checkCountEl = document.getElementById("check-count");

const grid = document.getElementById("chart-grid");
const linePath = document.getElementById("line-path");
const areaPath = document.getElementById("area-path");

let cachedTargets = [];
let cachedResults = [];

toggleFormBtn.addEventListener("click", () => {
  monitorForm.classList.toggle("hidden");
});

monitorForm.addEventListener("submit", async (event) => {
  event.preventDefault();
  feedback.textContent = "Salvando monitor...";

  const payload = {
    name: document.getElementById("monitor-name").value.trim(),
    url: document.getElementById("monitor-url").value.trim(),
    expectedStatus: 200,
    intervalMs: Number(document.getElementById("monitor-interval").value) * 1000,
    timeoutMs: Number(document.getElementById("monitor-timeout").value),
    retries: 1,
    isActive: true,
  };

  try {
    const response = await fetch(`${API_BASE}/targets`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    });

    if (!response.ok) {
      throw new Error(`Erro ao criar monitor (${response.status})`);
    }

    feedback.textContent = "Monitor criado com sucesso.";
    monitorForm.reset();
    monitorForm.classList.add("hidden");
    await refresh();
  } catch (error) {
    feedback.textContent = error.message;
  }
});

runChecksBtn.addEventListener("click", async () => {
  feedback.textContent = "Executando checks...";
  try {
    const response = await fetch(`${API_BASE}/checks/run`, { method: "POST" });
    if (!response.ok) {
      throw new Error(`Erro ao executar checks (${response.status})`);
    }
    feedback.textContent = "Checks executados. Atualizando painel...";
    setTimeout(refresh, 800);
  } catch (error) {
    feedback.textContent = error.message;
  }
});

async function refresh() {
  const [targetsResult, resultsResult] = await Promise.allSettled([
    fetch(`${API_BASE}/targets`),
    fetch(`${API_BASE}/results?limit=200`),
  ]);

  if (targetsResult.status === "fulfilled" && targetsResult.value.ok) {
    const body = await targetsResult.value.json();
    cachedTargets = body.data ?? [];
  }

  if (resultsResult.status === "fulfilled" && resultsResult.value.ok) {
    const body = await resultsResult.value.json();
    cachedResults = body.data ?? [];
  }

  renderTable(cachedTargets, cachedResults);
  renderChart(cachedResults);
  renderStats(cachedResults);
}

function renderTable(targets, results) {
  if (!targets.length) {
    table.innerHTML = `<div class="status-row"><span class="website">Nenhum monitor cadastrado</span><span>-</span><span>-</span><span>-</span></div>`;
    return;
  }

  const byTarget = groupByTarget(results);

  table.innerHTML = targets
    .map((target) => {
      const history = byTarget.get(target.id) ?? [];
      const latest = history[0];
      const uptime = computeUptime(history);
      const status = getStatus(latest);
      const responseTime = formatResponseTime(latest);

      return `
        <div class="status-row">
          <div class="website"><span class="dot"></span>${escapeHtml(target.url)}</div>
          <div class="status ${status.className}"><span class="dot" style="background:${status.color}"></span>${status.label}</div>
          <div>${responseTime}</div>
          <div>${uptime}</div>
        </div>
      `;
    })
    .join("");
}

function groupByTarget(results) {
  const map = new Map();

  const ordered = [...results].sort((a, b) => new Date(b.checkedAt) - new Date(a.checkedAt));
  ordered.forEach((item) => {
    if (!map.has(item.targetId)) {
      map.set(item.targetId, []);
    }
    map.get(item.targetId).push(item);
  });

  return map;
}

function computeUptime(history) {
  if (!history.length) {
    return "--";
  }

  const upCount = history.filter((entry) => entry.isUp).length;
  const uptime = (upCount / history.length) * 100;
  return `${uptime.toFixed(2)}%`;
}

function getStatus(result) {
  if (!result) {
    return { label: "Unknown", className: "slow", color: "#9aa7bb" };
  }

  if (!result.isUp) {
    return { label: "Offline", className: "offline", color: "#f44336" };
  }

  if (Number(result.responseTimeMs) > 700) {
    return { label: "Slow", className: "slow", color: "#f6bf26" };
  }

  return { label: "Online", className: "online", color: "#0f9d73" };
}

function formatResponseTime(result) {
  if (!result) {
    return "--";
  }

  if (!result.isUp && result.error) {
    return "Timeout";
  }

  return `${Number(result.responseTimeMs ?? 0)} ms`;
}

function renderStats(results) {
  const values = results.map((item) => Number(item.responseTimeMs || 0)).filter((x) => x >= 0);
  if (!values.length) {
    avgLatencyEl.textContent = "0 ms";
    maxLatencyEl.textContent = "0 ms";
    checkCountEl.textContent = "0";
    return;
  }

  const sum = values.reduce((acc, current) => acc + current, 0);
  const avg = Math.round(sum / values.length);
  const max = Math.max(...values);

  avgLatencyEl.textContent = `${avg} ms`;
  maxLatencyEl.textContent = `${max} ms`;
  checkCountEl.textContent = String(values.length);
}

function renderChart(results) {
  const points = [...results]
    .sort((a, b) => new Date(a.checkedAt) - new Date(b.checkedAt))
    .slice(-50)
    .map((entry) => Number(entry.responseTimeMs || 0));

  grid.innerHTML = "";
  [40, 80, 120, 160, 200].forEach((y) => {
    const line = document.createElementNS("http://www.w3.org/2000/svg", "line");
    line.setAttribute("x1", "0");
    line.setAttribute("x2", "1000");
    line.setAttribute("y1", String(y));
    line.setAttribute("y2", String(y));
    line.setAttribute("class", "grid-line");
    grid.appendChild(line);
  });

  if (points.length < 2) {
    linePath.setAttribute("d", "");
    areaPath.setAttribute("d", "");
    return;
  }

  const maxValue = Math.max(...points, 100);
  const stepX = 1000 / (points.length - 1);
  const coords = points.map((value, index) => {
    const x = index * stepX;
    const y = 220 - (value / maxValue) * 180;
    return [x, y];
  });

  const lineD = coords
    .map(([x, y], idx) => `${idx === 0 ? "M" : "L"}${x.toFixed(2)},${y.toFixed(2)}`)
    .join(" ");

  const areaD = `${lineD} L 1000,220 L 0,220 Z`;

  linePath.setAttribute("d", lineD);
  areaPath.setAttribute("d", areaD);
}

function escapeHtml(value) {
  return String(value)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#39;");
}

refresh();
setInterval(refresh, 20000);
