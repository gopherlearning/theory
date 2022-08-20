# time, context

## Список возможных обозначений в макете:
1, Jan, January — месяц;
2, Mon, Monday — день;
3, 15 — час в 12- и 24-часовом формате соответственно;
4 — минута;
5, .0-000000000, .9-999999999 — секунда и доли секунды;
06, 2006 — год;
-0700, -07:00, -07, Z0700, Z07:00, Z07 — часовой пояс;
pm, PM — время суток;
MST — аббревиатура часового пояса.

## Отдельные точки во времени можно сравнить методами:

// Equal проверяет, равны ли два момента времени.
func (t Time) Equal(u Time) bool

// After проверяет, наступил ли момент времени t после u.
func (t Time) After(u Time) bool

// Before проверяет, наступил ли момент времени t перед u.
func (t Time) Before(u Time) bool 

// вычислить интервал между временными точками
duration := stop.Sub(start)



https://gokit.io/faq/#dependency-injection-mdash-why-is-func-main-always-so-big



Перечислим ещё раз приведённые в этом уроке рекомендации — сохраните их, чтобы использовать в качестве чек-листа для самопроверки:
* Не создавайте интерфейсы заранее.
* Принимайте интерфейсы, возвращайте структуры.
* Не создавайте лишние указатели.
* Давайте пакетам понятные имена.
* Ставьте context.Context первым параметром функции.
* Не используйте select с одним выбором.
* Не добавляйте избыточные nil-проверки срезов.
* Не ставьте break в конструкции switch — case.
* Не используйте без необходимости пустой идентификатор.




func insertVideos(ctx context.Context, db *sql.DB, videos ...Video) error {
    // шаг 1 — объявляем транзакцию
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    // шаг 1.1 — если возникает ошибка, откатываем изменения
    defer tx.Rollback()

    // шаг 2 — готовим инструкцию
    stmt, err := tx.PrepareContext(ctx, "INSERT INTO videos(title, description, views, likes) VALUES(?,?,?,?)")
    if err != nil {
        return err
    }
    // шаг 2.1 — не забываем закрыть инструкцию, когда она больше не нужна
    defer stmt.Close()

    for _, v := range videos {
        // шаг 3 — указываем, что каждое видео будет добавлено в транзакцию
        if _, err = stmt.ExecContext(ctx, v.Title, v.Description, v.Views, v.Likes); err != nil {
            return err
        }
    }
    // шаг 4 — сохраняем изменения'nj
    return tx.Commit()
yt 




// плохо
if err == ErrAccessDenied {
}
// хорошо
if errors.Is(err, ErrAccessDenied) {
}

var myErr *MyError
// плохо
if myErr, ok = err.(*MyError); ok {
}
// хорошо
if errors.As(err, &myErr) {
} 