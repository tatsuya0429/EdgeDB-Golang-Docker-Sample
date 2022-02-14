module default {
  type User {
    required property username -> str {
      constraint exclusive;
    };
    required property password -> str;
  }

  type Todo {
    required property title -> str;
    property description -> str;
    property status -> str {
      constraint one_of('Todo', 'WIP', 'Done');
      default := 'Todo';
    }
    property deadline -> cal::local_date;
    required link user -> User;
  }
}