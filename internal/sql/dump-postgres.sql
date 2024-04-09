PGDMP     8    %        	        |            postgres    15.6    15.3                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    5    postgres    DATABASE     �   CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1254';
    DROP DATABASE postgres;
                postgres    false                       0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    3353                        2615    16398    test    SCHEMA        CREATE SCHEMA test;
    DROP SCHEMA test;
                postgres    false            �            1259    16418    task_status    TABLE     0  CREATE TABLE test.task_status (
    id bigint NOT NULL,
    task_statu character varying(50),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    is_deleted boolean DEFAULT false,
    create_user_id bigint,
    update_user_id bigint
);
    DROP TABLE test.task_status;
       test         heap    postgres    false    7            �            1259    16417    task_status_id_seq    SEQUENCE     �   ALTER TABLE test.task_status ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME test.task_status_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            test          postgres    false    221    7            �            1259    16400    tasks    TABLE     �  CREATE TABLE test.tasks (
    id bigint NOT NULL,
    title character varying(200),
    description text,
    status_id bigint,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    is_deleted boolean DEFAULT false,
    create_user_id bigint,
    update_user_id bigint,
    difficulty bigint,
    duration bigint DEFAULT 1
);
    DROP TABLE test.tasks;
       test         heap    postgres    false    7            �            1259    16399    tasks_id_seq    SEQUENCE     �   ALTER TABLE test.tasks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME test.tasks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            test          postgres    false    217    7            �            1259    16409    users    TABLE     �  CREATE TABLE test.users (
    id bigint NOT NULL,
    first_name character varying,
    last_name character varying,
    email character varying,
    password character varying,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    is_deleted boolean DEFAULT false,
    create_user_id bigint,
    update_user_id bigint,
    user_level bigint
);
    DROP TABLE test.users;
       test         heap    postgres    false    7            �            1259    16408    users_id_seq    SEQUENCE     �   ALTER TABLE test.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME test.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            test          postgres    false    219    7                      0    16418    task_status 
   TABLE DATA           w   COPY test.task_status (id, task_statu, created_at, updated_at, is_deleted, create_user_id, update_user_id) FROM stdin;
    test          postgres    false    221   I                 0    16400    tasks 
   TABLE DATA           �   COPY test.tasks (id, title, description, status_id, created_at, updated_at, is_deleted, create_user_id, update_user_id, difficulty, duration) FROM stdin;
    test          postgres    false    217   �                 0    16409    users 
   TABLE DATA           �   COPY test.users (id, first_name, last_name, email, password, created_at, updated_at, is_deleted, create_user_id, update_user_id, user_level) FROM stdin;
    test          postgres    false    219                     0    0    task_status_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('test.task_status_id_seq', 8, true);
          test          postgres    false    220                       0    0    tasks_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('test.tasks_id_seq', 11, true);
          test          postgres    false    216                       0    0    users_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('test.users_id_seq', 13, true);
          test          postgres    false    218            ~           2606    16433    task_status task_status_pk 
   CONSTRAINT     V   ALTER TABLE ONLY test.task_status
    ADD CONSTRAINT task_status_pk PRIMARY KEY (id);
 B   ALTER TABLE ONLY test.task_status DROP CONSTRAINT task_status_pk;
       test            postgres    false    221            |           2606    16440    users users_pk 
   CONSTRAINT     J   ALTER TABLE ONLY test.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);
 6   ALTER TABLE ONLY test.users DROP CONSTRAINT users_pk;
       test            postgres    false    219                       2606    16434    tasks fk_status_id    FK CONSTRAINT     u   ALTER TABLE ONLY test.tasks
    ADD CONSTRAINT fk_status_id FOREIGN KEY (status_id) REFERENCES test.task_status(id);
 :   ALTER TABLE ONLY test.tasks DROP CONSTRAINT fk_status_id;
       test          postgres    false    221    217    3198               �   x��λ
�@��z�)��"!{1�� X)��Y��q'd���E����>~�"��1K��La\V��8Զ*\�}�s��������o����f!��)�e�n���LW
�3cCiR7����	y%�B?�'�i�MQ*u��x,���HB�=WJ} 7ڊ�         	  x����j�0�k�S�Nh��u�TW^��@b�?'��d�A �3�)��u���ez���}�N�߷����aȐ'9y>Q�DJA�����ET�ļƜ9�m=���1o-�ڲ�7	W��cc����<b.Ŕ=�a.�M ��RI9�,T<w����(�����}�&��7��>�L,��qS��4NE�l�ѣ:0(I*q�\��M�A�T�t;b1f�΄ ��Ҽ��a.W��ߟ ����ۖ�:,�n۲���Ž��         �   x�m�=N1�:s�� #ۉ�8���q�	��"����@�ɒ�{Շ��������帞��Q��<���,���M��lO�c�j�����3��h�B��H�ؔ�z��*�i$�Ǝ����L���������;�9���p�\w�t���V�{&Ijh���9F��](.�3$��b��y1i0Ф?,T �~o������0O��IV     