package repositories

import "strconv"

func (b *BookRepository) createQuery() string {
	return "" +
		"	INSERT INTO book ( " +
		"		title,	" +
		"		isbn, " +
		"		about, " +
		"		edition, " +
		"		page_number, " +
		"		release_date, " +
		"		author_id " +
		"	) VALUES ( " +
		"	$1, $2, $3, $4, $5, $6, $7) " +
		"	RETURNING id"
}

func (b *BookRepository) updateQuery() string {
	return "" +
		"	UPDATE book SET " +
		" 		title = $1, " +
		" 		isbn = $2, " +
		" 		about = $3, " +
		" 		edition = $4, " +
		" 		page_number = $5, " +
		" 		release_date = $6, " +
		" 		author_id = $7 " +
		"	WHERE id = $8"
}

func (b *BookRepository) deleteQuery() string {
	return "" +
		"	DELETE FROM book WHERE id = $1"
}

func (b *BookRepository) getByIdQuery(id int64) string {
	return "" +
		"	SELECT " +
		"		id,	" +
		"		title,	" +
		"		isbn, " +
		"		about, " +
		"		edition, " +
		"		page_number, " +
		"		release_date, " +
		"		author_id " +
		"	FROM book " +
		"	WHERE id = " + strconv.FormatInt(id, 10)
}

func (b *BookRepository) getPagedQuery(limit, offset int) string {
	return "" +
		"	SELECT " +
		"		id,	" +
		"		title,	" +
		"		isbn, " +
		"		about, " +
		"		edition, " +
		"		page_number, " +
		"		release_date, " +
		"		author_id " +
		"	FROM book " +
		"	ORDER BY title " +
		"	LIMIT " + strconv.Itoa(limit) +
		"	OFFSET " + strconv.Itoa(offset)
}
