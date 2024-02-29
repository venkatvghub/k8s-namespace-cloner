import { useState } from 'react';
import { Button, Container, DialogActions, DialogContent, DialogTitle, FormControl, Input, InputLabel } from '@mui/material';

const UpdateDeploymentDialog = (props) => {
  const deployment = props.deployment;
  const [image, setImage] = useState('');

  return (
    <div>
      <DialogTitle style={{ backgroundColor: 'rgba(34, 25, 67, 1)', color: 'white' }}>
        Update Deployment {deployment.name} in namespace {deployment.namespace}
      </DialogTitle>
      <DialogContent>
        <Container>
          <>
            <FormControl fullWidth>
              <InputLabel>Image</InputLabel>
              <Input value={image} onChange={(e) => setImage(e.target.value)} />
            </FormControl>
          </>
        </Container>
      </DialogContent>
      <DialogActions>
        <Button
          variant="outlined"
          disabled={!image}
          onClick={(e) => props.onSubmit(e, image, deployment)}
        >
          Submit
        </Button>
      </DialogActions>
    </div>
  );
};

export default UpdateDeploymentDialog;
