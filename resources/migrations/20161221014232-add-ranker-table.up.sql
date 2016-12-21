CREATE TABLE ranker
(
    cd_ranker serial primary key,
    cd_simulado int references simulado(cd_simulado)
);