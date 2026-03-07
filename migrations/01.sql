BEGIN;

INSERT INTO monitor_targets (
    id,
    name,
    url,
    method,
    timeout_ms,
    expected_status,
    retries,
    retry_delay_ms,
    active
)
VALUES
    (uuid_generate_v4(), 'Google', 'https://www.google.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'GitHub', 'https://github.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'OpenAI', 'https://openai.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Cloudflare', 'https://www.cloudflare.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Wikipedia', 'https://www.wikipedia.org', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Stack Overflow', 'https://stackoverflow.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'MDN Web Docs', 'https://developer.mozilla.org', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Go.dev', 'https://go.dev', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Docker', 'https://www.docker.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Kubernetes', 'https://kubernetes.io', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'NPM', 'https://www.npmjs.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'PyPI', 'https://pypi.org', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Stripe', 'https://stripe.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Vercel', 'https://vercel.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Netlify', 'https://www.netlify.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Amazon', 'https://www.amazon.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Microsoft', 'https://www.microsoft.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Apple', 'https://www.apple.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'YouTube', 'https://www.youtube.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Netflix', 'https://www.netflix.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'BBC', 'https://www.bbc.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'The New York Times', 'https://www.nytimes.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Globo', 'https://www.globo.com', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'UOL', 'https://www.uol.com.br', 'GET', 3000, 200, 1, 500, true),
    (uuid_generate_v4(), 'Mercado Livre', 'https://www.mercadolivre.com.br', 'GET', 3000, 200, 1, 500, true);

COMMIT;