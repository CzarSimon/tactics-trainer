import React from 'react';
import { useHistory } from 'react-router';
import { useCreateNewProblemSet } from '../../hooks';
import { CreateProblemSetRequest } from '../../types';
import { NewProblemSetForm } from './components/NewProblemSetForm';

export function NewProblemSetContainer() {
  const history = useHistory();
  const createNewProblemSet = useCreateNewProblemSet();

  const onCancel = () => {
    history.goBack();
  };

  const onSubmit = async (req: CreateProblemSetRequest) => {
    await createNewProblemSet(req);
    history.push('/');
  };

  return <NewProblemSetForm onSubmit={onSubmit} onCancel={onCancel} />;
}
