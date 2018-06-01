<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>UDP bug 查询系统</title>
</head>

<body>

<form action="index.php" method="post">
    <div>
	项目:<select name="project"> 
	<option value="%">any</option>
<?php
   $host        = "host=10.186.18.21";
   $port        = "port=5432";
   $dbname      = "dbname=gitlabhq_production";
   $credentials = "user=readonly password=readonly";

   $conn = pg_connect( "$host $port $dbname $credentials"  );
   if(!$conn){
      echo "Error : Unable to open database\n";
   }

   $sql =<<<EOF
select n.name,n.name||'/'||p.name from projects p,namespaces n where p.namespace_id=n.id order by n.name||'/'||p.name;
EOF;
   $ret = pg_query($conn,$sql);
   if(!$ret){
      echo pg_last_error($db);
      exit;
   }
   while($row = pg_fetch_row($ret))
   {
      echo "<option value=\"{$row[0]}\">{$row[1]}</option>";
   }
   pg_close(db);
?>

	</select> 
	查询方式:<select name="scope">
	<option value="tp">标题和正文</option>
	<option value="t">只标题</option>
	</select>
	Status:<select name="issue_stat">
	<option value="%">any</option>
	<option value="closed">closed</option>
	<option value="opened">opened</option>
	</select>
        关键字:<input type="text" name="keyword"/>
        <input type="submit" value="查询"/>

    </div>

</form>
<table border="1" cellspacing="0" cellpadding="0">
    <tr>
        <td width="80">项目</td>
        <td width="50" >ID</td>
        <td width="300" >标题</td>
        <td width="50" >状态</td>
        <td width="200" >URL</td>
    </tr>
<?php
   $host        = "host=10.186.18.21";
   $port        = "port=5432";
   $dbname      = "dbname=gitlabhq_production";
   $credentials = "user=readonly password=readonly";

   $conn = pg_connect( "$host $port $dbname $credentials"  );
   if(!$conn){
      echo "Error : Unable to open database\n";
   }

   $sql_tp =<<<EOF
SELECT 
    n.name || '/' || p.name AS project,
    i.iid AS issue_id,
    i.title AS issue_title,
    i.state AS issue_state,
    'http://10.186.18.21/' || n.name || '/'
        || p.name
        || '/issues/'
        || i.iid AS url
FROM
    issues i,
    projects p,
    namespaces n
WHERE
    i.project_id = p.id
        AND p.namespace_id = n.id
        AND p.name LIKE $1
        AND i.state LIKE $2
        AND (i.title LIKE $3
        OR i.description LIKE $4)
ORDER BY i.iid;
EOF;
   $sql_t =<<<EOF
SELECT 
    n.name || '/' || p.name AS project,
    i.iid AS issue_id,
    i.title AS issue_title,
    i.state AS issue_state,
    'http://10.186.18.21/' || n.name || '/'
        || p.name
        || '/issues/'
        || i.iid AS url
FROM
    issues i,
    projects p,
    namespaces n
WHERE
    i.project_id = p.id
        AND p.namespace_id = n.id
        AND p.name LIKE $1
        AND i.state LIKE $2
        AND i.title LIKE $3
ORDER BY i.iid;
EOF;

   $keyword_wildcard=sprintf("%%%s%%",$_POST["keyword"]);
   echo "已为您使用如下条件搜索: <br>";
   echo "项目: ",$_POST["project"],"<br>关键字: ",$keyword_wildcard,"<br>State: ",$_POST["issue_stat"];
   if ($_POST["scope"]=="tp")
   {
      $ret = pg_prepare($conn, "query_tp",$sql_tp);
      $ret=pg_execute($conn,"query_tp",array($_POST["project"],$_POST["issue_stat"],$keyword_wildcard,$keyword_wildcard));
   }else{
      $ret = pg_prepare($conn, "query_t",$sql_t);
      $ret=pg_execute($conn,"query_t",array($_POST["project"],$_POST["issue_stat"],$keyword_wildcard));
   }
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
        <td><a href=\"{$row[4]}\" target=\"_blank\">{$row[4]}</a></td>
        
    </tr>";
   }
   pg_close($db);
?>
</table>
</body>
</html>

