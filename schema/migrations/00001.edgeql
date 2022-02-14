CREATE MIGRATION m1g4c2igfy7nzjhckq77btpyqkhojppw6klj27qb353per4dz73d5q
    ONTO initial
{
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY password -> std::str;
      CREATE REQUIRED PROPERTY username -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };
  CREATE TYPE default::Todo {
      CREATE LINK user -> default::User;
      CREATE PROPERTY deadline -> cal::local_date;
      CREATE PROPERTY description -> std::str;
      CREATE PROPERTY status -> std::str {
          SET default := 'Todo';
          CREATE CONSTRAINT std::one_of('Todo', 'WIP', 'Done');
      };
      CREATE REQUIRED PROPERTY title -> std::str;
  };
};
