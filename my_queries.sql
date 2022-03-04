select avg(rating) from "MY_TABLE";
select max(rating) from "MY_TABLE";
select sum(rating) from "MY_TABLE";
select upper(website) from "MY_TABLE" where name = 'Simpsons Restaurant';
select coordinates from "MY_TABLE" where name = 'Simpsons Restaurant';
--0101000020E610000026C808A870C4FEBFDD408177F23B4A40

--select name, ST_AsText(coordinates) from "MY_TABLE" 
--    where ST_DWithin(coordinates, '0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000.0);
--select name, ST_AsText(coordinates), 
--    ST_Distance(coordinates, ST_GeomFromText('POINT(-1.922959 52.468337)', 4326)) from "MY_TABLE" 
--    where ST_DWithin(coordinates, 
--                    ST_GeomFromText('POINT(-1.922959 52.468337)', 4326), 1000.0);

--select *, 
--    ST_Distance(coordinates, ST_GeomFromText('POINT(-1.922959 52.468337)', 4326)) from "MY_TABLE" 
--    where ST_DWithin(coordinates, 
--                    ST_GeomFromText('POINT(-1.922959 52.468337)', 4326), 1000.0);

--select ST_AsText(ST_
--                ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(45.0)));
--select ST_AsText(
--    ST_MakePolygon(ST_MakeLine(Array[
--        ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(45.0)) :: geometry,
--        ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(135.0)) :: geometry,
--        ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(225.0)) :: geometry,
--        ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(315.0)) :: geometry,
--        ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(45.0)) :: geometry
--        ])));
--select name from "MY_TABLE"
--    where ST_Within(
--        coordinates :: geometry,
--        ST_MakePolygon(ST_MakeLine(Array[
--            ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(45.0)) :: geometry,
--            ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(135.0)) :: geometry,
--            ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(225.0)) :: geometry,
--            ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(315.0)) :: geometry,
--            ST_Project('0101000020E610000026C808A870C4FEBFDD408177F23B4A40', 1000, radians(45.0)) :: geometry])));

select name from "MY_TABLE"
    where ST_Within(
        coordinates :: geometry,
        ST_MakePolygon(ST_MakeLine(Array[
            ST_Project(ST_GeomFromText('POINT(-1.922959 52.468337)',4326) , 1000, radians(45.0)) :: geometry,
            ST_Project(ST_GeomFromText('POINT(-1.922959 52.468337)',4326) , 1000, radians(135.0)) :: geometry,
            ST_Project(ST_GeomFromText('POINT(-1.922959 52.468337)',4326) , 1000, radians(225.0)) :: geometry,
            ST_Project(ST_GeomFromText('POINT(-1.922959 52.468337)',4326) , 1000, radians(315.0)) :: geometry,
            ST_Project(ST_GeomFromText('POINT(-1.922959 52.468337)',4326) , 1000, radians(45.0)) :: geometry])));
--SELECT ST_AsText(ST_Buffer(
--    ST_GeomFromText('POINT(-1.922959 52.468337)', 4326),
-- 50, 'quad_segs=1'));
--select name, ST_SRID(coordinates), ST_AsText(coordinates) from "MY_TABLE";
--select regexp_replace('John Doe', '(.*) (.*)','\2, \1');
--select regexp_replace(website, '', '/^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)/img') from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
-- select regexp_replace(website, '/^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)/img', '\1') as result from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
--select regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1') as result from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
 --select regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1') as result from "MY_TABLE";
--select regexp_substr(website, '/^(?:https?:\/\/)?(?:[^@\n]+@)?(?:www\.)?([^:\/\n?]+)/img') from "MY_TABLE" where name = 'Simpsons Restaurant';
-- update "MY_TABLE"
--    set website = regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1') where name = 'CELINE London New Bond Street Store';
--select * from "MY_TABLE" where name = 'CELINE London New Bond Street Store';
--update "MY_TABLE"
--   set website = regexp_replace(website, '^(?:.*://)?(?:www\.)?([^:/]*).*$', '\1');
--\echo Hello
--select website as domain, count(*) as NUM from "MY_TABLE" where website is not null group by website;
--select * from
--    (select website, count(*) as NUM from "MY_TABLE" group by website) as foo
--    where NUM > 1;
--select name as spot, num as noOfTimes from
--    (select name, website, count(*) as num from "MY_TABLE" where website is not null group by website) as foo
--    where num > 1;

--select name as spot from "MY_TABLE"
--    where website in
--        (select website from
--            (select website, count(*) as num from "MY_TABLE" where website is not null group by website) as foo
--        where num > 1);
