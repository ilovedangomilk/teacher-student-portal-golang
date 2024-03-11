# teacher-student-portal-golang

To run the teacher-student-portal:

In main directory (teacher-student-portal-golang), run:
- go run main.go

After running the server, 
Open another instance of terminal, 
Teacher can use make-request.py to update a student's attendance, each time the make-request.py is ran, it increases the attendance count of that student by one
Teacher can edit student ID to increase by changing the student ID in make-request.py

StudentID is being added via DBeaver on PostgreSQL Database

In order to view the changes being made, go to your preferred web browser and type:
localhost:8080/students/{studentID}

E.g., 
localhost:8080/students/1

