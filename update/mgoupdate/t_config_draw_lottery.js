conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_config_draw_lottery').drop();

db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a71b"),
    "sort" : 1.0,
    "type" : 2,
    "name" : "10",
    "number" : 10.0,
    "percent" : 10.0,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a71c"),
    "sort" : 2.0,
    "type" : 7.0,
    "name" : "10",
    "number" : 10.0,
    "percent" : 15.0,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a71d"),
    "sort" : 3.0,
    "type" : 3.0,
    "name" : "0.2",
    "number" : 0.2,
    "percent" : 5.0,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a71e"),
    "sort" : 4.0,
    "type" : 1.0,
    "name" : "500",
    "number" : 500.0,
    "percent" : 24.5,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a71f"),
    "sort" : 5.0,
    "type" : 101,
    "name" : "1",
    "number" : 1.0,
    "percent" : 10.0,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a720"),
    "sort" : 6.0,
    "type" : 3,
    "name" : "1",
    "number" : 1.0,
    "percent" : 0.5,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a721"),
    "sort" : 7.0,
    "type" : 1,
    "name" : "1000",
    "number" : 1000.0,
    "percent" : 20.0,
    "version" : 1.0
}
);
db.t_config_draw_lottery.insert(
{
    "_id" : ObjectId("587c4e74e36aeb81eb71a722"),
    "sort" : 8.0,
    "type" : 1,
    "name" : "5000",
    "number" : 5000.0,
    "percent" : 5.0,
    "version" : 1.0
}
);
