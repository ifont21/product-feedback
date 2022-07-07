package infrastructure

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	domain "github.com/ifont21/product-feedback-service/internal/domain/onboarding"
)

type BoardDbRepository struct {
	db *sql.DB
}

func NewBoardDBRepository(db *sql.DB) BoardDbRepository {
	return BoardDbRepository{db}
}

func (r BoardDbRepository) GetAll(name string) ([]domain.BoardDTO, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM pf_board")
	boards := make([]domain.BoardDTO, 0)

	if err != nil {
		return []domain.BoardDTO{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		var name string
		var description string
		err = rows.Scan(&id, &name, &description)

		if err != nil {
			return boards, err
		}

		board := domain.BoardDTO{}
		board.ID = id
		board.Name = name
		board.Description = description

		boards = append(boards, board)
	}

	err = rows.Err()

	if err != nil {
		return []domain.BoardDTO{}, err
	}

	return boards, nil
}

func (r BoardDbRepository) GetById(id uuid.UUID) (domain.BoardDetailDTO, error) {
	sqlStatement := `
	SELECT pf_board.id,
		pf_board.name, 
		pf_board.description,
		pf_project.title
    	FROM pf_board
    JOIN pf_project
    ON pf_project.id = pf_board.pf_project_id
    WHERE pf_board.id = $1
	`
	var board domain.BoardDetailDTO

	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&board.ID, &board.Name, &board.Description, &board.Project)

	if err == sql.ErrNoRows {
		return domain.BoardDetailDTO{}, errors.New("no rows were returned")
	}

	return board, nil
}

func (r BoardDbRepository) Save(aggregate domain.BoardAggregate) error {
	sqlStatement := `
	INSERT INTO pf_board (id, name, description, pf_project_id)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		sqlStatement,
		aggregate.Board.GetID(),
		aggregate.Board.GetTitle(),
		aggregate.Board.GetDescription(),
		aggregate.Board.GetProjectID().ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r BoardDbRepository) Update(aggregate domain.BoardAggregate) error {
	return nil
}
