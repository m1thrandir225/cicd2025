CREATE TABLE votes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    option_id uuid REFERENCES poll_options (id) ON DELETE CASCADE NOT NULL,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE,
    voted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL
);