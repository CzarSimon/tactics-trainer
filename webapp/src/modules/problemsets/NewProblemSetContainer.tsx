import React from 'react';
import { useQueryClient } from 'react-query';
import { useHistory } from 'react-router';
import * as api from '../../api';
import { CreateProblemSetRequest } from '../../types';
import { NewProblemSetForm } from './components/NewProblemSetForm';

export function NewProblemSetContainer() {
  const history = useHistory();
  const queryClient = useQueryClient();

  const onCancel = () => {
    history.goBack();
  };

  const onSubmit = async (req: CreateProblemSetRequest) => {
    await api.createProblemSet(req);
    queryClient.invalidateQueries('problem-sets');
    history.push('/');
  };

  return <NewProblemSetForm onSubmit={onSubmit} onCancel={onCancel} />;
}
