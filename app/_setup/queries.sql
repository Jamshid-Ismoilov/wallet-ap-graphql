------------------Queries-------------------

select 
	balance
from users
where user_id = 1
;


select 
	comment,
	amount
from incomes
where user_id = 1 and category_id = 3;

select 
	comment,
	amount
from outgoings
where user_id = 1 and category_id = 1;


select 
	o.amount,
	to_char(o.created_at,'HH24:MI') as timebirth,	
	c.name,
	o.comment
from outgoings as o
join categories as c on o.category_id = c.category_id 
where extract(day from o.created_at) = 06 and extract(month from o.created_at) = 11 and extract(year from o.created_at) = 2021
order by timebirth
;

select 
	o.amount,
	extract(day from o.created_at) as days,	
	c.name,
	o.comment
from outgoings as o
join categories as c on o.category_id = c.category_id 
where extract(month from o.created_at) = 11 and extract(year from o.created_at) = 2021
order by days
;

select 
	o.amount,
	o.created_at::time as time,	
	c.name,
	o.comment
from outgoings as o
join categories as c on o.category_id = c.category_id 
where o.created_at::date = $1
order by days
;

type StatisticsBody struct {
	Amount       float64 `json:"amount"`
	Time         string  `json:"time"`
	CategoryName string  `json:"categoryName"`
	Comment      *string `json:"comment"`
}
