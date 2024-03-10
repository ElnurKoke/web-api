package store

import (
	"github.com/ElnurKoke/web-api.git/internal/model"
)

func (s *Store) GetProjectById(id int) (model.Project, error) {
	query := "SELECT * FROM projects WHERE projects.id = $1;"
	row := s.db.QueryRow(query, id)
	var p model.Project
	err := row.Scan(&p.ID,
		&p.ProjectName,
		&p.Category,
		&p.ProjectType,
		&p.ReleaseYear,
		&p.AgeCategory,
		&p.Duration,
		&p.Director,
		&p.Producer)
	if err != nil {
		return model.Project{}, err
	}

	return p, nil
}

func (s *Store) GetAllProjects() ([]model.Project, error) {
	var projects []model.Project
	rows, err := s.db.Query("SELECT * FROM projects")
	if err != nil {
		return []model.Project{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Project
		err := rows.Scan(&p.ID,
			&p.ProjectName,
			&p.Category,
			&p.ProjectType,
			&p.ReleaseYear,
			&p.AgeCategory,
			&p.Duration,
			&p.Director,
			&p.Producer)
		if err != nil {
			return []model.Project{}, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (s *Store) UpdateProject(project model.Project) error {
	query := `UPDATE projects 
			SET project_name=?, category=?, project_type=?, release_year=?, 
			age_category=?, duration=?, director=?, producer=? 
			WHERE id=?`
	if _, err := s.db.Exec(query, project.ProjectName,
		project.Category, project.ProjectType,
		project.ReleaseYear, project.AgeCategory,
		project.Duration, project.Director,
		project.Producer, project.ID); err != nil {
		return err
	}
	return nil
}
