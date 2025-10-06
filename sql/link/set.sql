CREATE OR REPLACE FUNCTION set(
    p_key TEXT,
    p_value TEXT
)
RETURNS VOID
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO key_value_store (key, value)
    VALUES (p_key, p_value)
    ON CONFLICT (key)
    DO UPDATE SET 
        value = EXCLUDED.value,
        updated_at = NOW();
END;
$$;
