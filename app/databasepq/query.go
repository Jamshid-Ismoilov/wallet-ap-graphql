package databasepq

var SELECTNEW = `
select exists(select email from users where email = $1)
`

var SELECTUSERID = `
select user_id from users where email = $1 and password = $2 and firstname = $3 and lastname = $4`

var INSERTUSER = `
insert into users (firstname, lastname, email, password, balance) values ($1,$2,$3,$4, 0.00)
returning user_id`

var UPDATEUSER = `
update users 
set firstname = $1,
lastname = $2,
password = $4
where email = $3
returning user_id
`

var DELETEUSERBYID = `
delete from users where user_id = $1`

var CHECKUSER = `
select exists(select email from users where email = $1 and password = $2 and firstname = $3 and lastname = $4)`

var CHECKUSERBYIDANDEMAIL = `
select exists(select email from users where user_id = $1 and email = $2)
`
var ADDINCOMEDB = `
insert into incomes (amount, user_id, category_id) values ($2, $1, $3)
returning income_id
`

var UPDATECREATEDATINCOME = `
update incomes set created_at = $1::timestamp where income_id = $2`

var UPDATECOMMENTINCOME = `
update incomes set comment = $1 where income_id = $2
`

var ADDOUTGOINGDB = `
insert into outgoings (amount, user_id, category_id) values ($2, $1, $3)
returning outgoing_id
`

var UPDATECREATEDATOUTGOING = `
update outgoings set created_at = $1::timestamp where outgoing_id = $2`

var UPDATECOMMENTOUTGOING = `
update outgoings set comment = $1 where outgoing_id = $2
`
var SETBALANCE = `
update users set balance = $1 where user_id = $2`

var GETDAILYOUTGOINGS = `
select 
	o.amount,
	o.created_at::time as time,	
	c.name,
	o.comment
from outgoings as o
join categories as c on o.category_id = c.category_id 
where extract(day from o.created_at) = $1 and extract(month from o.created_at) = $2 and extract(year from o.created_at) = $3 and user_id = $4 
order by time
`

var GETDAILYINCOMES = `
select 
	i.amount,
	i.created_at::time as time,	
	c.name,
	i.comment
from incomes as i
join categories as c on i.category_id = c.category_id 
where extract(day from i.created_at) = $1 and extract(month from i.created_at) = $2 and extract(year from i.created_at) = $3 and user_id = $4
order by time
`
var GETMONTHLYOUTGOINGS = `
select 
	o.amount,
	o.created_at::time as time,	
	c.name,
	o.comment
from outgoings as o
join categories as c on o.category_id = c.category_id 
where extract(month from o.created_at) = $1 and extract(year from o.created_at) = $2 and user_id = $3 
order by time
`

var GETMONTHLYINCOMES = `
select 
	i.amount,
	i.created_at::time as time,	
	c.name,
	i.comment
from incomes as i
join categories as c on i.category_id = c.category_id 
where extract(month from i.created_at) = $1 and extract(year from i.created_at) = $2 and user_id = $3 
order by time
`
var GETSPENDINGBYCATEGORY = `
select 
	o.amount,
	o.created_at::time as time,	
	c.name,
	o.comment
from outgoings as o
join categories as c on o.category_id = c.category_id 
where o.user_id = $1 and o.category_id = $2  
order by time
`

var GETINCOMESBYCATEGORY = `
select 
	i.amount,
	i.created_at::time as time,	
	c.name,
	i.comment
from incomes as i
join categories as c on i.category_id = c.category_id 
where i.user_id = $1 and i.category_id = $2  
order by time
`

var GETBALANCE = `
select
	balance
from users
where user_id = $1 and email = $2`
