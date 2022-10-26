create table places (
                        id serial primary key,
                        type_place varchar(256),
                        order_place int,
                        geometry_type varchar(256),
                        coordinates float[]
)

create table property_place (
                                id serial primary key,
                                place_id int references places(id),
                                cat varchar(256),
                                visibility varchar(256),
                                zoom int,
                                order_place varchar(256),
                                color varchar(256),
                                color_code varchar(256),
                                show_title text,
                                map_title text,
                                exclude_place bool,
                                title text,
                                type_place varchar(256),
                                map_icon int,
                                url varchar(256),
                                pic varchar(256),
                                coordinates float[],
                                attractions int[],
                                visible bool,
                                active bool
)
