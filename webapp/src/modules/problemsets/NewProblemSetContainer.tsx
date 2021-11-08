import React from 'react';
import { useHistory } from 'react-router';
import log from '@czarsimon/remotelogger';
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
    const { id } = await createNewProblemSet(req);
    log.info(`created new problem set id=${id}`);
    history.push('/');
  };

  return <NewProblemSetForm onSubmit={onSubmit} onCancel={onCancel} />;
}
