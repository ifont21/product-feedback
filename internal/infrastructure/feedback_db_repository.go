// Queries
package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	feedback "github.com/ifont21/product-feedback-service/internal/domain/feedback"
	onboarding "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type FeedbackDbRepository struct {
	db *sql.DB
}

func NewfeedbackDBRepository(db *sql.DB) FeedbackDbRepository {
	return FeedbackDbRepository{db}
}

func (r FeedbackDbRepository) GetBoardPostsByCriteria(boardId uuid.UUID, sortBy string, filterByCategory string) ([]feedback.PostDTO, error) {
	sqlStatement := `
	SELECT pf_post.id,
		pf_post.title,
		pf_post.description,
	    pf_category.name as category_name,
	    pf_post.votes
    FROM pf_post
    JOIN pf_category
    ON pf_post.pf_category_id = pf_category.id
    WHERE pf_post.pf_board_id = $1
    ORDER BY pf_post.created_at desc
    `
	rows, err := r.db.Query(sqlStatement, boardId)

	if err != nil {
		return []feedback.PostDTO{}, err
	}

	posts := make([]feedback.PostDTO, 0)

	defer rows.Close()

	for rows.Next() {
		var post feedback.PostDTO
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Category,
			&post.Votes,
		)

		if err != nil {
			return []feedback.PostDTO{}, err
		}

		posts = append(posts, post)
	}

	err = rows.Err()

	if err != nil {
		return []feedback.PostDTO{}, err
	}

	return posts, nil
}

func (r FeedbackDbRepository) GetPostById(postId int64) (feedback.PostDetailsDTO, error) {
	sqlStatement := `
	SELECT pf_post.id,
	    pf_board.id,
		pf_board.name,
		pf_post.title,
		pf_post.description,
	    pf_category.name as category_name,
		pf_post.status,
	    pf_post.votes,
		pf_post.updated_at
    FROM pf_post
    JOIN pf_category
    ON pf_post.pf_category_id = pf_category.id
	JOIN pf_board
	ON pf_post.pf_board_id = pf_board.id
    WHERE pf_post.id = $1
    `
	var post feedback.PostDetailsDTO
	var board onboarding.BoardDTO

	row := r.db.QueryRow(sqlStatement, postId)
	err := row.Scan(&post.ID,
		&board.ID,
		&board.Name,
		&post.Title,
		&post.Description,
		&post.Category,
		&post.Status,
		&post.Votes,
		&post.LastUpdatedDate,
	)

	if err == sql.ErrNoRows {
		return feedback.PostDetailsDTO{}, errors.New("no rows were returned")
	}

	post.Board = board

	return post, nil
}

// Comands
func (r FeedbackDbRepository) CreatePost(aggregate feedback.PostAggregate) error {
	sqlStatement := `
	 INSERT	INTO pf_post(id, pf_board_id, title, description, status, pf_category_id)
	 VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		sqlStatement,
		aggregate.GetPost().ID,
		aggregate.GetPost().
			BoardID.
			GetID(),
		aggregate.GetPost().Title,
		aggregate.GetPost().Description,
		aggregate.GetPost().Status,
		aggregate.GetPost().Category.ID,
	)

	if err != nil {
		return err
	}

	return err
}

func (r FeedbackDbRepository) UpdatePost(aggregate feedback.PostAggregate) error {
	sqlStatement := `
	UPDATE pf_post
	SET title = $2,
	    description = $3,
		status = $4, 
		pf_category_id = $5
	WHERE pf_post.id = $1
	`

	_, err := r.db.Exec(
		sqlStatement,
		aggregate.GetPost().ID,
		aggregate.GetPost().Title,
		aggregate.GetPost().Description,
		aggregate.GetPost().Status,
		aggregate.GetPost().Category.ID,
	)

	if err != nil {
		return err
	}

	return err
}

func (r FeedbackDbRepository) UpVote(aggregate feedback.PostAggregate) error {
	sqlStatement := `UPDATE pf_post SET votes = votes + 1 WHERE pf_post.id = $1`
	_, err := r.db.Exec(sqlStatement, aggregate.GetPost().ID)

	if err != nil {
		return err
	}

	return err
}

func (r FeedbackDbRepository) DeletePost(postId int64) error {
	// Implement Delete
	return nil
}
