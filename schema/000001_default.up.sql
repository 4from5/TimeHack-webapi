CREATE TABLE users
(
    user_id       serial PRIMARY KEY UNIQUE,
    username      text UNIQUE NOT NULL,
    email         text UNIQUE,
    password_hash text        NOT NULL
);

CREATE TABLE categories
(
    category_id serial PRIMARY KEY UNIQUE,
    user_id     int REFERENCES users (user_id) NOT NULL,
    colour      text DEFAULT 'black'
);

CREATE TABLE tasks
(
    task_id     serial PRIMARY KEY UNIQUE,
    user_id     int REFERENCES users (user_id)          NOT NULL,
    category_id int REFERENCES categories (category_id) NOT NULL,
    title       text                                    NOT NULL,
    description text,
    deadline    timestamp,
    date_time   timestamp                               NOT NULL,
    priority    int DEFAULT 1
);

CREATE TABLE event
(
    event_id             serial PRIMARY KEY UNIQUE,
    user_id              int REFERENCES users (user_id)          NOT NULL,
    category_id          int REFERENCES categories (category_id) NOT NULL,
    title                text                                    NOT NULL,
    description          text,
    start_timestamp      timestamp                               NOT NULL,
    end_timestamp        timestamp                               NOT NULL,
    is_full_day          bool DEFAULT false,
    location             text,
    period_time_days     int  default 0,
    end_period_timestamp timestamp

);

CREATE TABLE notion
(
    notion_id    serial PRIMARY KEY UNIQUE,
    user_id      int REFERENCES users (user_id)          NOT NULL,
    title        text                                    NOT NULL,
    notion_text  text,
    category_id  int REFERENCES categories (category_id) NOT NULL,
    created_date timestamp                               NOT NULL,
    last_update  timestamp
)