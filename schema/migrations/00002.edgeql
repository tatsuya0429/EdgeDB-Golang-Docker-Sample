CREATE MIGRATION m1ckvdtioksg6yfgrdklnoat3wg2gvbygd4vimycq27otcwsh7nmma
    ONTO m1g4c2igfy7nzjhckq77btpyqkhojppw6klj27qb353per4dz73d5q
{
  ALTER TYPE default::Todo {
      ALTER LINK user {
          SET REQUIRED USING (SELECT
              default::User
          FILTER
              (.id = <std::uuid>'13593254-8cec-11ec-86ad-8b77027fbe60')
          );
      };
  };
};
