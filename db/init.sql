
USE todo;

CREATE TABLE ticket(
  id          BIGINT(20)  NOT NULL AUTO_INCREMENT,
  title       VARCHAR(40) NOT NULL,
  created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

INSERT INTO ticket(title) VALUES ("Make UI feel good");
INSERT INTO ticket(title) VALUES ("fix back end feel good");
