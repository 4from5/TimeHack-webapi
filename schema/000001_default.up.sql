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
    title       text                           NOT NULL,
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

CREATE TABLE events
(
    event_id             serial PRIMARY KEY UNIQUE,
    user_id              int REFERENCES users (user_id)          NOT NULL,
    category_id          int REFERENCES categories (category_id) NOT NULL,
    title                text                                    NOT NULL,
    description          text,
    start_timestamp      timestamp                               NOT NULL,
    end_timestamp        timestamp                               NOT NULL,
    is_full_day          bool     DEFAULT false,
    location             text,
    repeat_period        interval default null,
    end_period_timestamp timestamp

);

CREATE TABLE notions
(
    notion_id    serial PRIMARY KEY UNIQUE,
    user_id      int REFERENCES users (user_id)          NOT NULL,
    category_id  int REFERENCES categories (category_id) NOT NULL,
    title        text                                    NOT NULL,
    notion_text  text,
    created_date timestamp                               NOT NULL,
    last_update  timestamp
);

--- SEEDER

--                  USERS
INSERT
INTO users(username, email, password_hash)
VALUES ('Cockpit', 'cockpit@mail.ru',
        '01C057776D3CB24BF6546A1442D6E8BBD3CB90D95F128FA6FCDB64623651540E84D7BA44D00E4CFFD1E51ECE965082E875712141B253A1FE3C75F5F1701D4B26');

INSERT
INTO users(username, password_hash)
VALUES ('Bashmak',
        'EEEAEFD769DDD20E546C7986A2E1ACD4762A35C9F755FF36EC138A0C8EF1FFD0ECFBA8CCD4CCAF1BAA7A0783796B476602C1B0E37F8DAAFF9A54FA58A26C9110');

--                  Categories
INSERT
INTO categories(user_id, title, colour)
VALUES (1, 'Личное', 'blue'),
       (1, 'Учёба', 'red'),
       (1, 'Чилл', 'yellow'),
       (2, 'Моё', 'green'),
       (2, 'Уник', 'orange'),
       (2, 'Отдых', 'violet');

--                  Tasks


--                  Notions
INSERT INTO notions(user_id, category_id, title, notion_text, created_date, last_update)
VALUES (1, 1, 'Личный дневник',
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean dictum ultricies dapibus. Pellentesque est nibh, posuere imperdiet diam sit amet, tincidunt vehicula lacus. Donec sed mollis libero, nec hendrerit eros. Nullam bibendum aliquet efficitur. Nulla sed scelerisque dolor. Curabitur tempor sit amet mauris a blandit. Vivamus blandit tincidunt quam, vitae auctor eros viverra vitae. Sed dolor felis, venenatis vitae libero ac, blandit varius mi. Suspendisse pellentesque est id arcu consectetur semper. Ut euismod dapibus urna. Aenean sit amet tortor orci. Praesent vel ligula libero.',
        '25 Nov 22 00:29 MSK', '25 Nov 22 00:32 MSK'),
       (1, 1, 'Список кайфа', 'Пока списка нет но кайф точно есть',
        '20 Nov 22 13:43 MSK', '23 Nov 22 00:29 MSK'),
       (1, 2, 'Долги по учёбе', 'А долгов то нет мы же крутые а вы чо хотели',
        '10 Nov 21 11:43 MSK', '20 Nov 22 13:43 MSK'),
       (1, 2, 'Физика ненавижу', 'тут формулы типа какие-то',
        '10 Nov 21 11:43 MSK', '10 Nov 22 11:43 MSK'),
       (1, 3, 'Список чилла нереального', 'А чилла то и нет потому что бомонка душит',
        '1 Nov 22 11:54 MSK', '11 Nov 22 23:30 MSK'),
       (1, 3, 'С пацанами на карики', 'точно шашлычок нужен и лимонадик(пиво не пью)',
        '12 Jul 22 17:20 MSK', '12 Jul 22 17:20 MSK')
--        (1, 4,),
--        (1, 4,),
--        (1, 5,),
--        (1, 5,),
--        (1, 6,),
--        (1, 6,)

--                  Events

