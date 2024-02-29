import { useState } from 'react';
import { Button, Container, DialogActions, DialogContent, DialogTitle, FormControl, Input, InputLabel } from '@mui/material';

const CloneNamespaceDialog = (props) => {
  const namespaceName = props.namespaceName;
  const [targetNamespace, setTargetNamespace] = useState('');

  return (
    <div>
      <DialogTitle style={{ backgroundColor: 'rgba(34, 25, 67, 1)', color: 'white' }}>
        Clone namespace {namespaceName}
      </DialogTitle>
      <DialogContent>
        <Container>
          <>
            <FormControl fullWidth>
              <InputLabel>Target Namespace</InputLabel>
              <Input value={targetNamespace} onChange={(e) => setTargetNamespace(e.target.value)} />
            </FormControl>
          </>
        </Container>
      </DialogContent>
      <DialogActions>
        <Button
          variant="outlined"
          disabled={!targetNamespace}
          onClick={(e) => props.onCloneClick(e, targetNamespace, namespaceName)}
        >
          Clone
        </Button>
      </DialogActions>
    </div>
  );
};

export default CloneNamespaceDialog;
