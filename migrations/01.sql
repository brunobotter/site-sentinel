BEGIN;

WITH base_sites(name, url) AS (
    VALUES
        ('Google', 'https://www.google.com'),
        ('GitHub', 'https://github.com'),
        ('OpenAI', 'https://openai.com'),
        ('Cloudflare', 'https://www.cloudflare.com'),
        ('Wikipedia', 'https://www.wikipedia.org'),
        ('Stack Overflow', 'https://stackoverflow.com'),
        ('MDN Web Docs', 'https://developer.mozilla.org'),
        ('Go.dev', 'https://go.dev'),
        ('Docker', 'https://www.docker.com'),
        ('Kubernetes', 'https://kubernetes.io'),
        ('NPM', 'https://www.npmjs.com'),
        ('PyPI', 'https://pypi.org'),
        ('Stripe', 'https://stripe.com'),
        ('Vercel', 'https://vercel.com'),
        ('Netlify', 'https://www.netlify.com'),
        ('Amazon', 'https://www.amazon.com'),
        ('Microsoft', 'https://www.microsoft.com'),
        ('Apple', 'https://www.apple.com'),
        ('YouTube', 'https://www.youtube.com'),
        ('Netflix', 'https://www.netflix.com'),
        ('BBC', 'https://www.bbc.com'),
        ('The New York Times', 'https://www.nytimes.com'),
        ('Globo', 'https://www.globo.com'),
        ('UOL', 'https://www.uol.com.br'),
        ('Mercado Livre', 'https://www.mercadolivre.com.br'),
        ('Reddit', 'https://www.reddit.com'),
        ('LinkedIn', 'https://www.linkedin.com'),
        ('X', 'https://x.com'),
        ('Instagram', 'https://www.instagram.com'),
        ('Facebook', 'https://www.facebook.com'),
        ('Twitch', 'https://www.twitch.tv'),
        ('Discord', 'https://discord.com'),
        ('Notion', 'https://www.notion.so'),
        ('Figma', 'https://www.figma.com'),
        ('Canva', 'https://www.canva.com'),
        ('Trello', 'https://trello.com'),
        ('Asana', 'https://asana.com'),
        ('Jira', 'https://www.atlassian.com/software/jira'),
        ('Bitbucket', 'https://bitbucket.org'),
        ('GitLab', 'https://gitlab.com'),
        ('DigitalOcean', 'https://www.digitalocean.com'),
        ('Heroku', 'https://www.heroku.com'),
        ('Linode', 'https://www.linode.com'),
        ('AWS', 'https://aws.amazon.com'),
        ('Google Cloud', 'https://cloud.google.com'),
        ('Azure', 'https://azure.microsoft.com'),
        ('HashiCorp', 'https://www.hashicorp.com'),
        ('Postman', 'https://www.postman.com'),
        ('Slack', 'https://slack.com'),
        ('Zoom', 'https://zoom.us')
),
ranked_sites AS (
    SELECT row_number() OVER () AS site_idx, name, url
    FROM base_sites
),
seed_rows AS (
    SELECT
        gs AS idx,
        rs.name || ' #' || gs AS name,
        rs.url
    FROM generate_series(1, 2000) AS gs
    INNER JOIN ranked_sites rs
        ON rs.site_idx = ((gs - 1) % (SELECT COUNT(*) FROM ranked_sites)) + 1
)
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
SELECT
    uuid_generate_v4(),
    name,
    url,
    'GET',
    3000,
    200,
    1,
    500,
    true
FROM seed_rows
ORDER BY idx;

COMMIT;
