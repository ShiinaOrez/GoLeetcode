# Write your MySQL query statement below
SELECT p.FirstName, p.LastName, a.City, a.State
FROM
person p LEFT JOIN address a
ON
p.personid=a.personid
