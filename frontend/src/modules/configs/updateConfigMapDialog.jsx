import { useState } from 'react';
import { Button, Container, DialogActions, DialogContent, DialogTitle, FormControl, Input, InputLabel } from '@mui/material';

const UpdateConfigMapDialog = (props) => {
  const config = props.config;
  const [value, setValue] = useState(config.value);

  return (
    <div>
      <DialogTitle style={{ backgroundColor: 'rgba(34, 25, 67, 1)', color: 'white' }}>
        Update Config {config.key}
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
          disabled={!value || value === config.value}
          onClick={(e) => props.onSubmit(e, config.key, value)}
        >
          Submit
        </Button>
      </DialogActions>
    </div>
  );
};

export default UpdateConfigMapDialog;
