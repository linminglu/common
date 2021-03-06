conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_config_task_info').drop();

db.t_config_task_info.insert(
{
    "taskid" : 1,
    "cateid" : 1,
    "tasktype" : "all_game_count",
    "gameid" : 0,
    "title" : "完成5局比赛",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 1,
    "description" : "奖励1000金币。",
    "reward" : [
        {
            "type" : 1,
            "amount" : 1000
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 2,
    "cateid" : 1,
    "tasktype" : "all_game_count",
    "gameid" : 0,
    "title" : "完成10局比赛",
    "sort" : 0,
    "tasksum" : 10,
    "repeatsum" : 1,
    "description" : "奖励1000金币。",
    "reward" : [
        {
            "type" : 1,
            "amount" : 1000
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 3,
    "cateid" : 1,
    "tasktype" : "all_game_count",
    "gameid" : 0,
    "title" : "完成15局比赛",
    "sort" : 0,
    "tasksum" : 15,
    "repeatsum" : 1,
    "description" : "奖励1000金币。",
    "reward" : [
        {
            "type" : 1,
            "amount" : 1000
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 4,
    "cateid" : 1,
    "tasktype" : "all_game_count",
    "gameid" : 0,
    "title" : "完成20局比赛",
    "sort" : 0,
    "tasksum" : 20,
    "repeatsum" : 1,
    "description" : "奖励2000金币。",
    "reward" : [
        {
            "type" : 1,
            "amount" : 2000
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 5,
    "cateid" : 1,
    "tasktype" : "all_game_count",
    "gameid" : 0,
    "title" : "完成25局比赛",
    "sort" : 0,
    "tasksum" : 25,
    "repeatsum" : 1,
    "description" : "奖励2000金币。",
    "reward" : [
        {
            "type" : 1,
            "amount" : 2000
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 6,
    "cateid" : 1,
    "tasktype" : "all_game_count",
    "gameid" : 0,
    "title" : "完成50局比赛",
    "sort" : 0,
    "tasksum" : 50,
    "repeatsum" : 1,
    "description" : "奖励3000金币。",
    "reward" : [
        {
            "type" : 1,
            "amount" : 3000
        }
    ]
}
);

db.t_config_task_info.insert(
{
    "taskid" : 7,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.08
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 8,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 10,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.09
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 9,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 15,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.10
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 10,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 20,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.11
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 11,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 25,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 12,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 30,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.10
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 13,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 35,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.08
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 14,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 40,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.09
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 15,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 45,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.11
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 16,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 50,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 17,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 55,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.09
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 18,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.11
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 19,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 60,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.08
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 20,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 65,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 21,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 70,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.11
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 22,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 75,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.09
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 23,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 80,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.10
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 24,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 85,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.09
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 25,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 90,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.11
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 26,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 95,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.09
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 27,
    "cateid" : 2,
    "tasktype" : "all_game_win",
    "gameid" : 0,
    "title" : "赢5局比赛",
    "sort" : 0,
    "tasksum" : 100,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);

db.t_config_task_info.insert(
{
    "taskid" : 101,
    "cateid" : 3,
    "tasktype" : "daily_share",
    "gameid" : 0,
    "title" : "每日分享领三次红包",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 1,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
db.t_config_task_info.insert(
{
    "taskid" : 102,
    "cateid" : 3,
    "tasktype" : "daily_share",
    "gameid" : 0,
    "title" : "每日分享领三次红包",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 3,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);

//麻将红包任务
db.t_config_task_info.insert(
{
    "taskid" : 201,
    "cateid" : 4,
    "tasktype" : "mj_win_coin_count",
    "gameid" : 4,
    "title" : "麻将金币场红包任务",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 10,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
//斗地主红包任务
db.t_config_task_info.insert(
{
    "taskid" : 202,
    "cateid" : 4,
    "tasktype" : "ddz_win_coin_count",
    "gameid" : 3,
    "title" : "斗地主金币场红包任务",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 10,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
//炸金花红包任务
db.t_config_task_info.insert(
{
    "taskid" : 203,
    "cateid" : 4,
    "tasktype" : "zjh_win_coin_count",
    "gameid" : 4,
    "title" : "炸金花金币场红包任务",
    "sort" : 0,
    "tasksum" : 5,
    "repeatsum" : 10,
    "description" : "奖励0.08红包。",
    "reward" : [
        {
            "type" : 3,
            "amount" : 0.12
        }
    ]
}
);
