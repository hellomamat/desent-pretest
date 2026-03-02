package handlers

import (
	"desent-pretest/database"
	"desent-pretest/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year,omitempty"`
}

type UpdateBookRequest struct {
	Title  *string `json:"title,omitempty"`
	Author *string `json:"author,omitempty"`
	Year   *int    `json:"year,omitempty"`
}

// Level 3: Create Book
func CreateBook(c *fiber.Ctx) error {
	var req CreateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Level 7: Validation
	if req.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title is required",
		})
	}
	if req.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "author is required",
		})
	}

	book := models.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := database.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// Level 3: Get All Books (with Level 6: Search & Paginate)
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	query := database.DB.Model(&models.Book{})

	// Level 6: Search by author
	if author := c.Query("author"); author != "" {
		query = query.Where("author ILIKE ?", "%"+author+"%")
	}

	// Level 6: Search by title
	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// Level 6: Pagination
	page, _ := strconv.Atoi(c.Query("page", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "0"))

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Order("id ASC").Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch books",
		})
	}

	return c.JSON(books)
}

// Level 3: Get Book by ID
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		// Level 7: Not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "book not found",
		})
	}

	return c.JSON(book)
}

// Level 4: Update Book
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "book not found",
		})
	}

	var req UpdateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.Author != nil {
		book.Author = *req.Author
	}
	if req.Year != nil {
		book.Year = *req.Year
	}

	if err := database.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update book",
		})
	}

	return c.JSON(book)
}

// Level 4: Delete Book
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "book not found",
		})
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete book",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
