-- Create a table to manage lists
CREATE TABLE IF NOT EXISTS lists (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    priority INT NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL
);

-- Create a table to manage sub-lists
CREATE TABLE IF NOT EXISTS sub_lists (
    id SERIAL PRIMARY KEY,
    list_id INT REFERENCES lists(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    priority INT NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL
);

-- Create a table to manage uploaded files
CREATE TABLE IF NOT EXISTS uploaded_files (
    id SERIAL PRIMARY KEY,
    list_id INT REFERENCES lists(id) ON DELETE CASCADE,
    sub_list_id INT REFERENCES sub_lists(id) ON DELETE CASCADE,
    file_name VARCHAR(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL
);

-- Insert some dummy data into lists
INSERT INTO lists (title, description,priority) VALUES
    ('List 1', 'Description for List 1',1),
    ('List 2', 'Description for List 2',2);

-- Insert some dummy data into sub_lists
INSERT INTO sub_lists (list_id, title, description,priority) VALUES
    (1, 'Sub-List 1', 'Description for Sub-List 1 in List 1',1),
    (1, 'Sub-List 2', 'Description for Sub-List 2 in List 1',2),
    (2, 'Sub-List Aa', 'Description for Sub-List Aa in List 2',1);

-- Insert some dummy data into uploaded_files
INSERT INTO uploaded_files (list_id,sub_list_id, file_name) VALUES
    (1, 1, 'file1.txt'), 
    (2, 2, 'file2.pdf');
