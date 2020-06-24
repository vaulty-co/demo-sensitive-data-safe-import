<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link href="https://fonts.googleapis.com/css?family=Open+Sans|Raleway:800" rel="stylesheet">
<link rel="stylesheet" href="/css/normalize.css">
<link rel="stylesheet" href="/css/skeleton.css">
<style>
    body {
		font-family: 'Open Sans', sans-serif;
	}

	h4 {
		color: #000;
		font-family: 'Raleway', sans-serif;
		font-weight: 800;
		font-size: 30px;
		margin: 0;		
	}	

	.container {
		padding-top: 50px;
		text-align: left;
	}

</style>
</head>
<body>
	<div class="container">
		<h4>Users</h4>
		<table class="pure-table pure-table-horizontal u-full-width">
			<thead>
				<tr>
					<th>Name</th>
					<th>Email</th>
					<th>Phone</th>
					<th>IP</th>
					<th>Card</th>
					<th>SSN</th>
				</tr>
			</thead>
			<tbody>
				{{range .Users}}
				<tr>
					<td>{{.Name}}</td>
					<td>{{.Email}}</td>
					<td>{{.PhoneNumber}}</td>
					<td>{{.IP}}</td>
					<td>{{.CardNumber}}</td>
					<td>{{.SSN}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</div>
</body>
</html>
