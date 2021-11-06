-- if trigger has an error, how to delete trigger:  drop trigger trigger_name on table_name;

-------------------------------------------------------------------------------------
-- this is a trigger function than updates users balance
-- after inserting into INCOMES table
create or replace function income_func () returns trigger language plpgsql as 
	$$	
	declare
		newBalance decimal(18,2) := 0;
	begin 
		select balance from users into newBalance where user_id = new.user_id;

		newBalance = newBalance + new.amount;

		update users set balance = newBalance where user_id = new.user_id;

		return null;
	end;

	$$;


create trigger income_trigger
after insert on incomes
for each row execute procedure income_func();

-------------------------------------------------------------------------------------
-- this is a trigger function than updates users balance
-- after inserting into OUTGOINGS table
create or replace function outgoings_func () returns trigger language plpgsql as 
	$$	
	declare
		newBalance decimal(18,2) := 0;
	begin 
		select balance from users into newBalance where user_id = new.user_id;

		newBalance = newBalance - new.amount;

		update users set balance = newBalance where user_id = new.user_id;

		return null;
	end;

	$$;


create trigger outgoing_trigger
after insert on outgoings
for each row execute procedure outgoings_func();
