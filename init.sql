drop table places cascade;

create table places (
                        id serial    primary key,
                        qr 			 varchar(256),
                        order_place  int,
                        coordinates  float[],
                        facilities   text[],
                        video_route  varchar(256),
                        tickets_link varchar(256),
                        tickets_text text,
                        cat          varchar(256),
                        visibility   varchar(256),
                        color        varchar(256),
                        color_code   varchar(256),
                        preview_text text,
                        detail_text  text,
                        title        varchar(256),
                        time_work    varchar(256),
                        type_place   varchar(256),
                        url          varchar(256),
                        pic          varchar(256),
                        code         varchar(256),
                        is_event	 bool
);

create table place_event_sharing (
                        id serial primary key,
                        place_id int references places(id),
                        event_id int references places(id)
);
