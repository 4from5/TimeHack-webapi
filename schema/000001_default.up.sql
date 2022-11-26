CREATE TABLE users
(
    user_id       serial PRIMARY KEY UNIQUE,
    username      text UNIQUE NOT NULL,
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
    task_id       serial PRIMARY KEY UNIQUE,
    user_id       int REFERENCES users (user_id)          NOT NULL,
    category_id   int REFERENCES categories (category_id) NOT NULL,
    title         text                                    NOT NULL,
    description   text,
    deadline      timestamp,
    date_time     timestamp                               NOT NULL,
    creation_date timestamp                               NOT NULL,
    priority      int DEFAULT 1
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
    is_full_day          bool DEFAULT false,
    event_location       text,
    repeat_period_days   int  default null,
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
INTO users(username, password_hash)
VALUES ('Cockpit',
        '73616c745f666f725f68617368a536a2c57d148f488a7b214a07356710a331f1256a6d39d647fee4c2da52cdc7511afe0aa59bddfc14bd3844c9cdb4ef54f049c1453a65b75c8634c5ab056837');

INSERT
INTO users(username, password_hash)
VALUES ('Bashmak1',
        '73616c745f666f725f68617368c1708d426030b4d70f5b7fc5b21475e2da1fa1a2959f3616346209566c69f03fa02d129cded76c4836de53a1cd376f84be188cff6a735ab25eb61476eb932de6');

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
INSERT INTO tasks(user_id, category_id, title, description, deadline, date_time, creation_date, priority)
VALUES (1, 1, 'Встретиться с кентом', 'ну это с Тёмиком короче пересечься бы', '30 Nov 22',
        '25 Nov 22', '25 Nov 22 12:02 MSK', 2),
       (1, 1, 'Жёстко заняться саморазвитием', '', '1 Jan 70',
        '27 Nov 22', '27 Nov 22 11:10 MSK', 3),
       (1, 2, 'Курсач по ААСОИУ', 'Шуку надо чот написать так и не понял тип того', '29 Dec 22',
        '27 Nov 22', '12 Sep 22 12:00', 1),
       (1, 2, 'ДЗ Тервер', '', '29 Nov 22',
        '28 Nov 22', '27 Nov 22 11:37 MSK', 3),
       (1, 3, 'Жёстко зачилить с девушкой', '', '1 Jan 70',
        '30 Nov 22', '30 Nov 22', 2),
       (1, 3, 'Жёстко с парнями зачилить', 'в баньку сгонять', '1 Jan 70',
        '1 Dec 22', '30 Nov 22', 3);


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
        '12 Jul 22 17:20 MSK', '12 Jul 22 17:20 MSK');

--                  Events
INSERT INTO events (user_id, category_id, title, description, start_timestamp, end_timestamp, is_full_day,
                    event_location, repeat_period_days, end_period_timestamp)
VALUES (1, 1, 'Спортзал', 'В качалочку погнал братик', '26 Nov 22 20:00 MSK',
        '26 Nov 22 22:00 MSK', false, 'на семёновской крч', 7, '31 Dec 22 00:00 MSK'),
       (1, 2, 'Сделать аасоиу', 'там курсач', '26 Nov 22 20:00 MSK',
        '26 Nov 22 20:00 MSK', true, '.', 0, '1 Jan 70 00:01 MSK'),
       (1, 3, 'Жёсткий сон', '.', '27 Nov 22 00:00 MSK',
        '27 Nov 22 06:00 MSK', true, 'в кроватке', 1, '30 Dec 22 00:00 MSK');

