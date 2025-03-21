-- name: GetCategoryByID :one
SELECT * FROM categories WHERE id = ? LIMIT 1;

-- name: GetSubCategoriesByID :many
SELECT * FROM categories WHERE parent_id = ?;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: CreateCategory :execresult
INSERT INTO categories (name, parent_id) VALUES (?, ?);

-- name: UpdateCategory :execresult
UPDATE categories SET name = ?, parent_id = ? WHERE id = ?;

-- name: DeleteCategory :execresult
DELETE FROM categories WHERE id = ?;
