INSERT INTO usuarios (nome, nick, email, senha)
VALUES 

('Usuario 1', 'usuario_1', 'usuario1@gmail.com', '$2a$10$jUtPklDod/O.UagCw4L5IO681fiJflVLxqgwXmqacaMAifaN1YyX2 ');
('Usuario 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$jUtPklDod/O.UagCw4L5IO681fiJflVLxqgwXmqacaMAifaN1YyX2 ');
('Usuario 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$jUtPklDod/O.UagCw4L5IO681fiJflVLxqgwXmqacaMAifaN1YyX2 ');

INSERT INTO seguidores(usuarios_id, seguidor_id)
VALUES

(1, 2),
(3, 1),
(1, 3);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicação do usuario 1", "esse é a publicação do usuario 1!", 1),
("Publicação do usuario 2", "esse é a publicação do usuario 2!", 2),
("Publicação do usuario 3", "esse é a publicação do usuario 3!", 3);
