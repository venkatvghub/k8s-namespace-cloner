import { useState } from 'react';
import { Button, Container, DialogActions, DialogContent, DialogTitle, FormControl, Input, InputLabel } from '@mui/material';

const UpdateSecretDialog = (props) => {
  const secret = props.secret;
  const [value, setValue] = useState('');

  return (
    <div>
      <DialogTitle style={{ backgroundColor: 'rgba(34, 25, 67, 1)', color: 'white' }}>
        Update Secret {secret.key}
      </DialogTitle>
      <DialogContent>
        <Container>
          <>
            <FormControl fullWidth>
              <InputLabel>Value</InputLabel>
              <Input value={value} onChange={(e) => setValue(e.target.value)} />
            </FormControl>
          </>
        </Container>
      </DialogContent>
      <DialogActions>
        <Button
          variant="outlined"
          disabled={!value}
          onClick={(e) => props.onSubmit(e, secret.key, value)}
        >
          Submit
        </Button>
      </DialogActions>
    </div>
  );
};

export default UpdateSecretDialog;
