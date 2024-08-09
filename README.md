# Ghidra Auth

## Setup

Compile / download the ghidra auth binary. Edit the following lines in `server.conf`

``` apacheconf
wrapper.app.parameter.1=-a4
wrapper.app.parameter.2=-jaas ./jaas.conf
wrapper.app.parameter.3=-u
```

Then modify `jaas.conf`
``` apacheconf
auth {
	ghidra.server.security.loginmodule.ExternalProgramLoginModule REQUIRED
		PROGRAM="/path/to/ghidraAuth"
		TIMEOUT="10000"
		ARG_00="/path/to/ghidra/server/"
		ARG_01="/path/to/ghidra/repositories/~admin/"
		ARG_02="default_repository_name" // default repository
		ARG_03="https://example.com/auth"
	;
};
```

Note: it sends a json document with `username` and `password` fields as a post request to the API. If the response is `ok` the user will be authorized. It's beyond the scope of this to actually handle that part, maybe in the future :3
 
