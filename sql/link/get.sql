CREATE OR REPLACE FUNCTION get(
    p_key TEXT
)
RETURNS TABLE(value_text TEXT)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
    SELECT kv.value FROM key_value_store kv WHERE kv.key = p_key;
END;
$$;
