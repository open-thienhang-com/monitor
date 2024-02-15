INSERT INTO campaign (time_start, time_end, quality, description, status) VALUES
(now(), now() + interval '1days', 1, 'This is a test discount', 'New'),
(now() + interval' 1 days', now() + interval '2 day', 2, 'This is another discount', 'In progress'),
(now() + interval '2 day', now() + interval '3 day', 3, 'This is a third test discount', 'Completed'),
(now() + interval '3 day', now() + interval '4 day', 1, 'This is a fourth test discount', 'New'),
(now() + interval '4 day', now() + interval '5 day', 2, 'This is a fifth test discount', 'In progress'),
(now() + interval '5 day', now() + interval '6 day', 3, 'This is a sixth test discount', 'Completed'),
(now() + interval '6 day', now() + interval '7 day', 1, 'This is a seventh test discount', 'New'),
(now() + interval '7 day', now() + interval '8 day', 2, 'This is an eighth test discount', 'In progress'),
(now() + interval '8 day', now() + interval '9 day', 3, 'This is a ninth test discount', 'Completed'),
(now() + interval '9 day', now() + interval '10 day', 1, 'This is a tenth test discount', 'New')
