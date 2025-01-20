-- Создает таблицу books с минимальным набором полей
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,        -- Уникальный идентификатор книги
                       title VARCHAR(255) NOT NULL,  -- Название книги
                       author VARCHAR(255) NOT NULL, -- Автор книги
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Время создания записи
);
