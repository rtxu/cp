# Write your MySQL query statement below


SELECT Department, Employee, Salary
FROM (
    SELECT d.Name as Department, e.Name as Employee, e.Salary #, IFNULL(d2.ThirdSalary, 0) as ThirdSalary
    FROM Employee e
    LEFT JOIN Department d ON e.DepartmentId = d.Id
    LEFT JOIN (
        SELECT DISTINCT DepartmentId as Id, Salary as ThirdSalary
        FROM (
            SELECT DepartmentId, 
                DENSE_RANK() OVER(PARTITION BY DepartmentId ORDER BY Salary DESC) as SalaryRank,
                Salary
            FROM Employee
        ) a
        WHERE SalaryRank = 3
    ) d2 ON e.DepartmentId = d2.Id
) a
WHERE Salary >= ThirdSalary AND Department IS NOT NULL



