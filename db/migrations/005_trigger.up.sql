
CREATE OR REPLACE FUNCTION synch_balance() RETURNS TRIGGER AS $$

DECLARE
  account_from_id UUID;
  account_to_id UUID;
  amount INTEGER;

BEGIN

account_from_id := NEW.account_from_id;
account_to_id := NEW.account_to_id;
amount := NEW.amount;

UPDATE accounts SET balance = balance - amount WHERE id = account_from_id;
UPDATE accounts SET balance = balance + amount WHERE id = account_to_id;

 RETURN NULL;
END; 
$$ LANGUAGE plpgsql;


CREATE TRIGGER synch_balance_trigger AFTER INSERT ON transactions FOR EACH ROW EXECUTE FUNCTION update_balance();
