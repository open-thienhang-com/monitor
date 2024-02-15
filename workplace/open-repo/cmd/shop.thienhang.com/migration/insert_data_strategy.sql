INSERT INTO strategy (discount_type,time_start, time_end, status) VALUES
('discount 1',now(), now() + interval '1days', 'New'),
('discount 2',now() + interval' 1 days', now() + interval '2 day',  'In progress'),
('discount 3',now() + interval '2 day', now() + interval '3 day',  'Completed'),
('discount 4',now() + interval '3 day', now() + interval '4 day',  'New'),
('discount 5',now() + interval '4 day', now() + interval '5 day', 'In progress'),
('discount 6',now() + interval '5 day', now() + interval '6 day',  'Completed'),
('discount 7',now() + interval '6 day', now() + interval '7 day',  'New'),
('discount 8',now() + interval '7 day', now() + interval '8 day',  'In progress'),
('discount 9',now() + interval '8 day', now() + interval '9 day',  'Completed'),
('discount 10',now() + interval '9 day', now() + interval '10 day', 'New')
