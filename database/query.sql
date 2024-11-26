-- Active: 1732632262439@@127.0.0.1@5432@olympics@public
INSERT INTO exercises (title, description, image_path)
VALUES ('Mounten pose', 'super simple mounte pose', 'mounten_pose_yoga.png');

INSERT INTO exercise_parameter (exercise_id, type, value)
VALUES (1, 'duration', '30');

INSERT INTO exercise_parameter (exercise_id, type, value)
VALUES (1, 'times', '1');

INSERT INTO exercise_parameter (exercise_id, type, value)
VALUES (1, 'times', '10');