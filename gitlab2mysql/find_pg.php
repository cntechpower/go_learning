<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>UDP bug 查询系统</title>
</head>

<body>

<form action="find_pg.php" method="post">
    <div>
	项目
	<select name="project"> 
	<option value="%">any</option> 
	<option value="umc">umc</option> 
	<option value="ucore">ucore</option> 
	<option value="uagent">uagent</option> 
	<option value="udeploy">udeploy</option> 
	<option value="umon">umon</option> 
	<option value="urds">urds</option> 
	<option value="ustats">ustats</option> 
	<option value="bug-firehole">bug-firehole</option> 
	<option value="ushard">ushard</option> 
	<option value="pm-work-log">pm-work-log</option> 
	<option value="uguard">uguard</option> 
	<option value="ulogstash">usql</option> 
	<option value="pm-work-log">pm-work-log</option> 
	</select> 
        关键字<input type="text" name="name"/>
        <input type="submit" value="查询"/>

    </div>

</form>
<table border="1" cellspacing="0" cellpadding="0">
    <tr>
        <td width="200">项目</td>
        <td width="200" >issue ID</td>
        <td width="200" >标题</td>
        <td width="200" >URL</td>
    </tr>
<?php
   $host        = "host=10.186.18.21";
   $port        = "port=5432";
   $dbname      = "dbname=gitlabhq_production";
   $credentials = "user=readonly password=readonly";

   $db = pg_connect( "$host $port $dbname $credentials"  );
   if(!$db){
      echo "Error : Unable to open database\n";
   }

   $sql =<<<EOF
SELECT
    p.name AS project_name,
    i.iid AS issue_id,
    i.title AS issue_title,
    'http://10.186.18.21/universe/'
        || p.name
        || '/issues/'
        || i.iid AS url
FROM
    issues i,
    projects p
WHERE
    i.project_id = p.id AND p.name like $1
        AND (i.title LIKE $2
        OR i.description LIKE $3);;

EOF;

   $ret = pg_prepare($db, "query_all",$sql);
   $keyword_wildcard=sprintf("%%%s%%",$_POST["name"]);
   echo "已为您使用如下条件搜索: <br>";
   echo "项目: ",$_POST["project"],"<br>关键字: ",$keyword_wildcard;
   $ret=pg_execute($db,"query_all",array($_POST["project"],$keyword_wildcard,$keyword_wildcard));
   if(!$ret){
      echo pg_last_error($db);
      exit;
   }
   while($row = pg_fetch_row($ret))
   {
        echo " <tr>
        <td>{$row[0]}</td> 
        <td>{$row[1]}</td>
        <td>{$row[2]}</td>
        <td>{$row[3]}</td>
        
    </tr>";
   }
   pg_close($db);
?>
</table>
</body>
</html>

