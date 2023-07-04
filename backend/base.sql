CREATE SCHEMA public;

CREATE TABLE public.user(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    email VARCHAR(30) UNIQUE,
    password VARCHAR(30) NOT NULL
);

CREATE TABLE public.todo_list(
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT  todo_list_fk FOREIGN KEY (user_id) REFERENCES public.user(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE public.todo(
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    describtion VARCHAR NOT NULL,
    is_ready BOOLEAN DEFAULT false,
    list_id INT,
    CONSTRAINT todo_fk FOREIGN KEY (list_id) REFERENCES public.todo_list(id) ON DELETE CASCADE ON UPDATE CASCADE 
);