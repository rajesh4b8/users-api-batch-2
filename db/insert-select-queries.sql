-- insert query
insert into users(first_name, last_name, email) values ('Rajesh', 'Reddy', 'rajeshtech4b8@gmail.com');

-- select query
select user_id, first_name, last_name, email, date_created from users where user_id = 1;