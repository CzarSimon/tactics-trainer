import { httpclient } from './httpclient';
import { CreateProblemSetRequest, Cycle, ProblemSet, Puzzle } from '../types';
import { wrapAndLogError } from './util';

const PUZZLE_SERVER_URL = '/api/puzzle-server';

export async function getPuzzle(id: string): Promise<Puzzle> {
  const { body, error, metadata } = await httpclient.get<Puzzle>({ url: `${PUZZLE_SERVER_URL}/v1/puzzles/${id}` });

  if (!body) {
    throw wrapAndLogError(`failed to fetch puzzle(id=${id})`, error, metadata);
  }

  return body;
}

export async function getProblemSet(id: string): Promise<ProblemSet> {
  const { body, error, metadata } = await httpclient.get<ProblemSet>({
    url: `${PUZZLE_SERVER_URL}/v1/problem-sets/${id}`,
  });

  if (!body) {
    throw wrapAndLogError(`failed to fetch problem sets id=${id}`, error, metadata);
  }

  return body;
}

export async function getProblemSets(): Promise<ProblemSet[]> {
  const { body, error, metadata } = await httpclient.get<ProblemSet[]>({
    url: `${PUZZLE_SERVER_URL}/v1/problem-sets`,
  });

  if (!body) {
    throw wrapAndLogError('failed to fetch problem sets', error, metadata);
  }

  return body;
}

export async function createProblemSet(req: CreateProblemSetRequest): Promise<ProblemSet> {
  const { body, error, metadata } = await httpclient.post<ProblemSet>({
    url: `${PUZZLE_SERVER_URL}/v1/problem-sets`,
    body: req,
  });

  if (!body) {
    throw wrapAndLogError('failed to create problem set', error, metadata);
  }

  return body;
}

export async function archiveProblemSet(id: string): Promise<void> {
  const { body, error, metadata } = await httpclient.delete<ProblemSet>({
    url: `${PUZZLE_SERVER_URL}/v1/problem-sets/${id}`,
  });

  if (!body) {
    throw wrapAndLogError(`failed to archive problemSet(id=${id})`, error, metadata);
  }
}

export async function listProblemSetCycles(id: string): Promise<Cycle[]> {
  const { body, error, metadata } = await httpclient.get<Cycle[]>({
    url: `${PUZZLE_SERVER_URL}/v1/problem-sets/${id}/cycles`,
  });

  if (!body) {
    throw wrapAndLogError(`failed to fetch cycles for problemSet(id=${id})`, error, metadata);
  }

  return body;
}

export async function createProblemSetCycle(id: string): Promise<Cycle> {
  const { body, error, metadata } = await httpclient.post<Cycle>({
    url: `${PUZZLE_SERVER_URL}/v1/problem-sets/${id}/cycles`,
  });

  if (!body) {
    throw wrapAndLogError(`failed to create cycle for problemSet(id=${id})`, error, metadata);
  }

  return body;
}

export async function getCycle(id: string): Promise<Cycle> {
  const { body, error, metadata } = await httpclient.get<Cycle>({
    url: `${PUZZLE_SERVER_URL}/v1/cycles/${id}`,
  });

  if (!body) {
    throw wrapAndLogError(`failed to fetch cycle(id=${id})`, error, metadata);
  }

  return body;
}

export async function updateCycle(id: string): Promise<Cycle> {
  const { body, error, metadata } = await httpclient.put<Cycle>({
    url: `${PUZZLE_SERVER_URL}/v1/cycles/${id}`,
  });

  if (!body) {
    throw wrapAndLogError(`failed to update cycle(id=${id})`, error, metadata);
  }

  return body;
}

export function getRandomPuzzleID(): string {
  const ids = [
    '649bbb7f-dc2c-47a2-bce0-c89dcac69b16',
    '0deddb83-f8e7-4cb3-8a1c-94dcd6f4bc94',
    '73205470-e281-4da8-85f1-2756969a4bc8',
    'bd8ffced-eff0-4ae0-b732-c1eda208a448',
    'ed61e756-e2b5-439f-b823-e249ed977128',
    'f3fce155-05db-4992-a46c-2078862b7385',
    'b6d35bd6-72b0-4ca8-9134-88eda6a45b47',
    'd58ca880-2685-4512-994a-426e731cc784',
    '8cdc2ea3-e810-4917-bfc4-75cd59f5ad72',
    '02019b85-7d72-4beb-89ee-8ada1691246e',
    '0eaf0e3c-0255-482b-aabf-7d3e60b15b7d',
    '3d32485b-3a02-424c-bd47-6123583c4e6f',
    '5e5e5f39-e0b4-46b5-921a-753b08c51cf9',
    'a663a3f5-3011-46e5-bb4b-1cd6ec5be5f1',
    '9e0795c9-1842-4ef0-9032-79acd2fa222a',
    '4934819a-88dc-4c4b-8f1f-b1732ab9b95f',
    '4632e2e3-20aa-42db-a950-dd21f719e867',
    'bfd535bf-efd1-48e5-967a-c74ce21bc6f2',
    '0bfb3be7-3a30-44e7-8db9-5134795aae84',
    'e2458f63-130f-4451-9b38-3e611611879e',
    '2f5157f0-8c46-4b28-9a4f-d3a65fc99ad3',
    '0676f68d-3504-4827-9319-1faad443a357',
    '727330c2-7ade-457a-8b63-89b835370690',
    '171e43bf-73ab-43ad-b43d-690fd09ee991',
    '5cb6fa90-0929-4428-8ade-1e47febc4d52',
    '5e5db1f6-a11f-41dc-bf04-eee156d9457b',
    '9d231b35-275a-47dd-a269-427977f8e6fc',
    '74ab7955-81c7-47cb-b2a4-165f15b0225b',
    '4e1168d0-7186-4a60-be99-0126d754b613',
    '3b79e629-aaca-40ce-a955-30860e39db1e',
    'cf51564e-8fab-40c2-9ee3-4fb8784b8322',
    '207bad4f-2ca2-4068-8bb9-79002e1e78f7',
    '9dc9ad40-fc2b-4bd3-8249-b063feeac7aa',
    'ac8571dc-c90d-414c-883a-d51c919659bb',
    'aa088a08-1846-49e8-af19-ebc7d09daed4',
    '47a205d1-3383-49b5-8bcc-ac68d6ab4bb0',
    'fef2e49e-f4fe-4e16-b1ac-c2eadedb8fbc',
    'bb116654-6c0b-4860-8c71-8139a74f8730',
    '3c66c99e-fb75-4ef9-8ded-f8acd1e94233',
    '2ba4f12a-ca4c-48ee-a65a-d57f29dc578e',
    'ac2a8d8d-59bb-4631-833a-a3a6da24bef4',
    '9b075673-83e3-4852-a4dc-60c85d847e16',
    '9475c745-3556-4324-8c2c-d65d7041d942',
    'cb3fd1e5-d056-423b-84d2-fa686dafd5dc',
    'a83b8a6e-b3dd-4da6-bb73-a3866977b950',
    '3b0ce14b-437e-4a24-9530-f72ddff56aae',
    'f8644dc3-19f2-4f97-8ccf-29b85d5369d5',
    '9d7f70c1-7d74-4d72-b158-b1ef2b706cb6',
    '1b53dca9-af09-4774-9468-477c02eae6fa',
    'ac16b30d-ddf1-44c8-903c-2aae051901f5',
    '8b89156e-4dfa-4322-8b7b-a06bcf9ad0fb',
    '632b0d8b-fd4f-4eac-937d-b24dc5084684',
    '5f635914-bb0b-4a8c-84ed-d2496d06997a',
    'd08f81d9-0ef9-4d6b-aac0-f1eb370e9a80',
    '9f4ae348-2332-49a9-9d64-94fb149b2add',
    'a3020b9e-ae23-43d0-a4a6-e92ef678d2f5',
    '88b45aa1-ea7f-49c8-8c4c-da984320b2df',
    '55c1f65b-5133-4ae1-a479-889889c5638d',
    '314e809b-73e8-45a9-a0a4-ebc43e6022e5',
    'b99c8367-bbbd-47c8-800d-fb27f053967d',
    'f87d3ac7-6204-421b-b40d-8d7b5237404f',
    'cedaf89a-bc5e-4df0-9524-4404cfdf510f',
    'ddd55a0e-59a8-45b1-977f-edc8519e1b4c',
    'ee8ee41a-97ad-4cf8-9281-fb528242ab16',
    '2102130d-483a-4cee-896b-70b3165760bd',
    'a9b3046d-54b4-4d91-b5a3-e54d7d2f3c12',
    '92171989-94d1-4cef-89dc-a614cf1a02e2',
    'f0a6b7d5-18b3-4542-9dac-baf4f880c7b8',
    '9c36b05e-a2c0-4ff5-96d2-97cf9b031f75',
    '68eaed19-e051-401d-ab3b-59ed0600c93f',
    '00aeff1e-f890-411d-b0d5-671926447629',
    '9c58534f-ff77-4e9c-9fd9-93ccdee115ae',
    'e611fb32-4795-463c-b340-b037ef4146dc',
    'fa6e806c-1ee9-4235-b76a-80efa93ff43b',
    '3a6b2763-7977-420a-9414-0af21695ed19',
    '553f2150-5174-4a83-8093-f81890efcf6e',
    'b940ab9d-db01-45f0-afff-6aaf2655c0f1',
    '877ff116-cf0a-455e-9db6-5ed651b48928',
    '7cd88ace-3400-44de-98af-f2a8916c7cfd',
    'e79f72a2-c7da-4503-9b95-41faa5e14206',
    '288bcc46-799d-4411-9538-7d7a499be355',
    '4e1db5a4-2564-44a2-bb9f-888360d2b59c',
    'ab2d67dd-5fd2-4104-be1b-160ae3fd4f74',
    '627ec265-dd4b-4632-bb4d-b30bd2ecb1cf',
    'b9db7876-ce9c-426f-a0d9-1588dd255250',
    'd80dce92-6ae2-4bba-8e2f-aa8eb5e0720f',
    '6148433f-144b-4e13-ba11-0a1eade16fe7',
    'b56c8acf-463a-45b7-879a-5055bcd32f0a',
    '68dfd80c-3582-4438-b587-56d544a89abb',
    'c8fa165f-806f-4ee2-94ce-39d17437db9e',
    '676a7463-5d1b-4b84-992f-7308376ea352',
    '491e750d-0915-4bb9-9157-278cd79bab8c',
    '0f5f718e-d685-40c5-9a1b-952e88e05505',
    'e98f3cf5-b9a9-4706-b248-dac903fdfc1b',
    'e22167e4-cb86-41b3-bc0a-6ffa777aae6c',
    '784699c7-7eea-4909-ab0d-517b69119f3e',
    'bbd93998-1a33-4156-8d6f-f56d78d25610',
    '09dda912-581a-4e78-96ce-ff1dc5eaa22b',
    'b7870697-ecc0-4fc7-ac9d-6277ff52073c',
    'c0d43a4d-be8b-4c54-88b7-1d763607131c',
    '6ae3f1fd-51d7-487a-99f0-5b7e72496ff0',
    '161e8451-bf80-41b7-96d6-1c698c80d8f3',
    '22b6fb05-434d-4f97-ae7d-2e379373fd3c',
    'c9892ac7-c06e-4088-8eb1-d59a3f2bc8d5',
    'e8893b7c-6acf-4acd-9788-170803979da3',
    'dff4365f-e55b-416b-9102-faa9ca039171',
    'bf000605-9278-47d8-ad86-38b0e171129a',
    '4ed2aafc-0f50-49d3-892a-245ccca2d3f4',
    'ee40b08f-5800-4d4e-974e-d8c173cc84f4',
    '9f6b1357-65f3-4e61-b36f-cf5894d609b7',
    '56d48dfe-82de-4d0c-abe3-e6132416c35b',
    '3a5eda20-12e5-48ea-bfe8-ae3f06d5238b',
    'fa447cd2-132b-4048-82ec-ecc07c31d925',
    '18a5d9e9-9325-4923-b5dd-98546a7701ef',
    '5602f9a6-478c-4f67-bbb4-77058ebb02d3',
    '541562f7-a89e-4ed5-b6dc-d8403e620dd5',
    'eb0c6a79-a402-41ca-9411-db41591a6956',
    '879cd883-f05a-473c-ad8c-463840496f8e',
    '95e065de-5e30-43ee-9bbd-60e871fcebec',
    '333e8cfd-7918-4cd3-8b79-8d9fc2025000',
    '35c3455e-86cb-4899-85f6-d419434d0475',
    'ced27e90-20bc-4cf2-be88-cbba0ebfcb8a',
    '0ab5a8d5-cd7d-4bf5-a178-5cb8d9612518',
    'bb918928-568d-4ff8-a168-8ea3a6c5c901',
    '934c1603-accb-4da3-9124-aa43c30df259',
    '9b9cdad9-280c-48b4-99b4-58387ce12e38',
    '2fd29c2f-1454-4606-99a8-738665496fca',
    '651bfb4f-de59-4c41-946b-4c8a17118552',
    '4ac9bc08-7e01-4c43-af31-a7f7b4ff119d',
    'f7273099-5c02-4ce4-8750-187b36ac2e4d',
    '118fd515-b073-436e-9554-9e63eb14da0e',
    'f3838a8a-9c07-495b-b425-d0980f0354af',
    '713b17b2-1073-4c4e-bb47-1dd0acd16452',
    'b96bdb72-e024-44a2-bd38-35a96a1cb05b',
    '541e575f-bbf6-4f33-9ac0-eb424cb971d7',
    '14a24212-311f-4e5e-aaab-e0bb266cfa0d',
    'd34aeb3e-852c-405a-ba95-6bb8f0ec0fce',
    'c6b30064-3897-4a75-abf6-2a5b4f91e8f6',
    '8ac02c3f-a6e0-4146-8cf2-d483a9715b5d',
    'f9aaaa4a-613c-494f-a256-31dc99d34f8a',
    'f085a788-f360-469a-b800-0a7d1b499913',
    '08624e5b-a527-415b-8007-b1bf94d20362',
    'bb414b6b-7cf5-4c69-bc8c-f116a7a64cb8',
    '25b9d578-3bb8-470d-86ae-9ebca9fd149c',
    '52bcbf65-fad2-4bf4-a4de-3bb55a3a29f8',
    '0c9e2ec5-08bc-4074-9308-28d037fb6496',
    'b5c66066-5ee3-4d55-9feb-aad28993f451',
    '379edbb8-3422-4a9a-b7d1-683f76836c5f',
    '3f8b4a93-097d-4d39-996b-8057a65d3a2c',
    '83ca5961-efbd-4b12-9ebc-a77b8db13aa1',
    'f85ad42c-a824-4813-8d6a-327ccf5675e2',
    '9a3aa452-1fda-4f3e-bc92-916f7550f096',
    '656bbf15-6524-4070-a815-f30492b75a9b',
    '4f18b52d-87f7-4098-971f-11cb0957c9e2',
    '14dfd736-ac35-4faf-ad2c-4db0d8a447c9',
    '3df175d3-e6ac-4af3-a133-e4d727e81737',
    '7079665b-7fe0-468a-8351-0b6c4d321ee7',
    '3efec48d-b817-4849-a1a9-9b696387b845',
    '8642e88e-9c9d-48a9-9008-903668cd3242',
    '421ae340-5f88-4b70-aee5-345c727b3c81',
    'c40ffb38-bc37-40d0-a148-54c944b3299c',
    '5e235b31-53da-41ae-bbe8-7b7a3ab6c9be',
    '9278306b-ed8c-42fa-b1cf-f1632eebfd0c',
    '661b6a4d-b053-4e28-95d8-81c787411add',
    'c0e389eb-9162-4a24-aa70-d9c312a9397d',
    'f05a122d-5ac6-4a99-b95d-7b3f54855e77',
    '3210551f-e8cf-4a1a-8eca-6e16be6c11d7',
    '723b9447-ebab-4d89-af63-949aa971998f',
    '7b76e934-0f65-4653-aab8-0293aa13f72e',
    '51e40b86-a7b7-4f58-80c1-64c04d12a973',
    'c8ff08e1-c45e-4e2a-a3d1-603e9f143680',
    '11d40b67-e20d-4bd7-95d2-b8801858c9bd',
    'a827cc06-e480-46b6-b5ce-ee7e8ce8c03a',
    'fb2ad6fd-cf26-442e-9ee7-4845bbefeb93',
    '49b8a8fd-9f5c-437b-9d98-38fdf962f8b5',
    '2fcce75b-3910-4e69-b40f-90392823698e',
    'b89a3b1d-0297-4c79-b5ef-f68e677de9ac',
    '7508f217-f617-49de-a396-2e5201665b93',
    'b911acec-e345-49fc-8a51-25162e1d1e30',
    '6b5fca59-da47-42a6-8a9c-380795683cad',
    '6150667c-b464-48a2-b2d9-5e90f744c23b',
    '100472cb-e0a9-47b5-8d98-0de2b6528373',
    'f6835286-3448-4ebe-8c11-624d272026cb',
    '716e15de-6969-4816-afa6-6fa084cb486e',
    'ca4c8c48-9a57-44a0-950c-d8e26142bfc1',
    '7266e434-d1ba-408b-8144-4a68140edf33',
    '8366284f-decf-4581-bef4-c344d2aa0a4e',
    'b5cc2e84-68d4-41f5-b01f-a34c40ac9d2b',
    '9e002e44-48ca-48be-b9f5-c05ce5021dda',
    '548aa129-ab49-482e-ada5-ff3c9e0df458',
    'ac28c8c4-dfa2-47dc-ac83-7e9ad144b76b',
    'd3f06cdc-14f6-4afd-885f-910e1cace3a4',
    'c79afb92-3d43-4591-b3c3-1bf570902ff0',
    'ed247adc-a130-4146-ac1f-5ff41cbf8592',
    '67cef9ba-c8e8-46c3-955c-823888b29ae9',
    'f0b193dc-8e9e-4cb3-a87c-1ee46a88c06b',
    '9e8331ee-4e9d-480b-af42-ad18cfbf5ddf',
    'd353182b-579d-492f-9a58-0144ffe047e0',
    '9fcb0f9f-6d79-4ed0-8e5c-1d63f6c1e65a',
    '5f1f5a2a-5f71-4ba7-bc67-bcd0e4e27a4e',
  ];

  const randomIdx = Math.floor(Math.random() * ids.length);
  return ids[randomIdx];
}
