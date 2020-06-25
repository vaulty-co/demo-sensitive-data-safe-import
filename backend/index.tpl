<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link href="/css/style.css" rel="stylesheet">
<link href="https://fonts.googleapis.com/css?family=Open+Sans:400|Raleway:800" rel="stylesheet">
</head>
<body>
	<main>
		<h4>Users</h4>
		<table>
			<thead>
				<tr>
					<th>Name</th>
					<th>Email</th>
					<th>Phone</th>
					<th>IP</th>
					<th>Card</th>
					<th>SSN</th>
					<th>Access Token</th>
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
					<td>{{.AccessToken}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
</main>
</body>
</html>
