import gql from 'graphql-tag';
import * as Urql from 'urql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Cursor: any;
  Metadata: any;
  Time: any;
};

export enum AccessLevel {
  Private = 'PRIVATE',
  Public = 'PUBLIC',
  View = 'VIEW'
}

export type CreateInstanceInput = {
  architecture: Scalars['String'];
  availability_zone: Scalars['String'];
  image_id: Scalars['String'];
  instance_id: Scalars['String'];
  instance_type: Scalars['String'];
  private_dns_name: Scalars['String'];
  private_ip_address: Scalars['String'];
  public_dns_name: Scalars['String'];
  public_ip_address: Scalars['String'];
  slide_id: Scalars['ID'];
  user_id: Scalars['ID'];
};

export type CreateSlideInput = {
  access_level?: InputMaybe<AccessLevel>;
  deleted?: InputMaybe<Scalars['Boolean']>;
  instance_id?: InputMaybe<Scalars['ID']>;
  name: Scalars['String'];
  path_token?: InputMaybe<Array<Scalars['String']>>;
  shared_with: Array<Scalars['String']>;
  size?: InputMaybe<Scalars['Int']>;
  user_id: Scalars['ID'];
};

export type CreateSlideWithTextInput = {
  access_level?: InputMaybe<AccessLevel>;
  name: Scalars['String'];
  shared_with?: InputMaybe<Array<Scalars['String']>>;
  user_id: Scalars['ID'];
};

export type CreateUserInput = {
  avatar_url?: InputMaybe<Scalars['String']>;
  bio?: InputMaybe<Scalars['String']>;
  email: Scalars['String'];
  full_name?: InputMaybe<Scalars['String']>;
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Instance = {
  __typename?: 'Instance';
  architecture: Scalars['String'];
  availability_zone: Scalars['String'];
  created_at: Scalars['Time'];
  id: Scalars['ID'];
  image_id: Scalars['String'];
  instance_id: Scalars['String'];
  instance_type: Scalars['String'];
  private_dns_name: Scalars['String'];
  private_ip_address: Scalars['String'];
  public_dns_name: Scalars['String'];
  public_ip_address: Scalars['String'];
  updated_at: Scalars['Time'];
};

export type InstanceConnection = {
  __typename?: 'InstanceConnection';
  edges: Array<InstanceEdge>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type InstanceEdge = {
  __typename?: 'InstanceEdge';
  cursor: Scalars['Cursor'];
  node: Instance;
};

export type InstanceOrder = {
  direction: OrderDirection;
  field?: InputMaybe<InstanceOrderField>;
};

export enum InstanceOrderField {
  CreatedAt = 'CREATED_AT',
  UpdatedAt = 'UPDATED_AT'
}

/**
 * InstanceWhereInput is used for filtering Instance objects.
 * Input was generated by ent.
 */
export type InstanceWhereInput = {
  and?: InputMaybe<Array<InstanceWhereInput>>;
  /** architecture field predicates */
  architecture?: InputMaybe<Scalars['String']>;
  architectureContains?: InputMaybe<Scalars['String']>;
  architectureContainsFold?: InputMaybe<Scalars['String']>;
  architectureEqualFold?: InputMaybe<Scalars['String']>;
  architectureGT?: InputMaybe<Scalars['String']>;
  architectureGTE?: InputMaybe<Scalars['String']>;
  architectureHasPrefix?: InputMaybe<Scalars['String']>;
  architectureHasSuffix?: InputMaybe<Scalars['String']>;
  architectureIn?: InputMaybe<Array<Scalars['String']>>;
  architectureLT?: InputMaybe<Scalars['String']>;
  architectureLTE?: InputMaybe<Scalars['String']>;
  architectureNEQ?: InputMaybe<Scalars['String']>;
  architectureNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** availability_zone field predicates */
  availabilityZone?: InputMaybe<Scalars['String']>;
  availabilityZoneContains?: InputMaybe<Scalars['String']>;
  availabilityZoneContainsFold?: InputMaybe<Scalars['String']>;
  availabilityZoneEqualFold?: InputMaybe<Scalars['String']>;
  availabilityZoneGT?: InputMaybe<Scalars['String']>;
  availabilityZoneGTE?: InputMaybe<Scalars['String']>;
  availabilityZoneHasPrefix?: InputMaybe<Scalars['String']>;
  availabilityZoneHasSuffix?: InputMaybe<Scalars['String']>;
  availabilityZoneIn?: InputMaybe<Array<Scalars['String']>>;
  availabilityZoneLT?: InputMaybe<Scalars['String']>;
  availabilityZoneLTE?: InputMaybe<Scalars['String']>;
  availabilityZoneNEQ?: InputMaybe<Scalars['String']>;
  availabilityZoneNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']>;
  createdAtGT?: InputMaybe<Scalars['Time']>;
  createdAtGTE?: InputMaybe<Scalars['Time']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']>>;
  createdAtLT?: InputMaybe<Scalars['Time']>;
  createdAtLTE?: InputMaybe<Scalars['Time']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']>>;
  /** slide edge predicates */
  hasSlide?: InputMaybe<Scalars['Boolean']>;
  hasSlideWith?: InputMaybe<Array<SlideWhereInput>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']>;
  idGT?: InputMaybe<Scalars['ID']>;
  idGTE?: InputMaybe<Scalars['ID']>;
  idIn?: InputMaybe<Array<Scalars['ID']>>;
  idLT?: InputMaybe<Scalars['ID']>;
  idLTE?: InputMaybe<Scalars['ID']>;
  idNEQ?: InputMaybe<Scalars['ID']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']>>;
  /** image_id field predicates */
  imageID?: InputMaybe<Scalars['String']>;
  imageIDContains?: InputMaybe<Scalars['String']>;
  imageIDContainsFold?: InputMaybe<Scalars['String']>;
  imageIDEqualFold?: InputMaybe<Scalars['String']>;
  imageIDGT?: InputMaybe<Scalars['String']>;
  imageIDGTE?: InputMaybe<Scalars['String']>;
  imageIDHasPrefix?: InputMaybe<Scalars['String']>;
  imageIDHasSuffix?: InputMaybe<Scalars['String']>;
  imageIDIn?: InputMaybe<Array<Scalars['String']>>;
  imageIDLT?: InputMaybe<Scalars['String']>;
  imageIDLTE?: InputMaybe<Scalars['String']>;
  imageIDNEQ?: InputMaybe<Scalars['String']>;
  imageIDNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** instance_id field predicates */
  instanceID?: InputMaybe<Scalars['String']>;
  instanceIDContains?: InputMaybe<Scalars['String']>;
  instanceIDContainsFold?: InputMaybe<Scalars['String']>;
  instanceIDEqualFold?: InputMaybe<Scalars['String']>;
  instanceIDGT?: InputMaybe<Scalars['String']>;
  instanceIDGTE?: InputMaybe<Scalars['String']>;
  instanceIDHasPrefix?: InputMaybe<Scalars['String']>;
  instanceIDHasSuffix?: InputMaybe<Scalars['String']>;
  instanceIDIn?: InputMaybe<Array<Scalars['String']>>;
  instanceIDLT?: InputMaybe<Scalars['String']>;
  instanceIDLTE?: InputMaybe<Scalars['String']>;
  instanceIDNEQ?: InputMaybe<Scalars['String']>;
  instanceIDNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** instance_type field predicates */
  instanceType?: InputMaybe<Scalars['String']>;
  instanceTypeContains?: InputMaybe<Scalars['String']>;
  instanceTypeContainsFold?: InputMaybe<Scalars['String']>;
  instanceTypeEqualFold?: InputMaybe<Scalars['String']>;
  instanceTypeGT?: InputMaybe<Scalars['String']>;
  instanceTypeGTE?: InputMaybe<Scalars['String']>;
  instanceTypeHasPrefix?: InputMaybe<Scalars['String']>;
  instanceTypeHasSuffix?: InputMaybe<Scalars['String']>;
  instanceTypeIn?: InputMaybe<Array<Scalars['String']>>;
  instanceTypeLT?: InputMaybe<Scalars['String']>;
  instanceTypeLTE?: InputMaybe<Scalars['String']>;
  instanceTypeNEQ?: InputMaybe<Scalars['String']>;
  instanceTypeNotIn?: InputMaybe<Array<Scalars['String']>>;
  not?: InputMaybe<InstanceWhereInput>;
  or?: InputMaybe<Array<InstanceWhereInput>>;
  /** private_dns_name field predicates */
  privateDNSName?: InputMaybe<Scalars['String']>;
  privateDNSNameContains?: InputMaybe<Scalars['String']>;
  privateDNSNameContainsFold?: InputMaybe<Scalars['String']>;
  privateDNSNameEqualFold?: InputMaybe<Scalars['String']>;
  privateDNSNameGT?: InputMaybe<Scalars['String']>;
  privateDNSNameGTE?: InputMaybe<Scalars['String']>;
  privateDNSNameHasPrefix?: InputMaybe<Scalars['String']>;
  privateDNSNameHasSuffix?: InputMaybe<Scalars['String']>;
  privateDNSNameIn?: InputMaybe<Array<Scalars['String']>>;
  privateDNSNameLT?: InputMaybe<Scalars['String']>;
  privateDNSNameLTE?: InputMaybe<Scalars['String']>;
  privateDNSNameNEQ?: InputMaybe<Scalars['String']>;
  privateDNSNameNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** private_ip_address field predicates */
  privateIPAddress?: InputMaybe<Scalars['String']>;
  privateIPAddressContains?: InputMaybe<Scalars['String']>;
  privateIPAddressContainsFold?: InputMaybe<Scalars['String']>;
  privateIPAddressEqualFold?: InputMaybe<Scalars['String']>;
  privateIPAddressGT?: InputMaybe<Scalars['String']>;
  privateIPAddressGTE?: InputMaybe<Scalars['String']>;
  privateIPAddressHasPrefix?: InputMaybe<Scalars['String']>;
  privateIPAddressHasSuffix?: InputMaybe<Scalars['String']>;
  privateIPAddressIn?: InputMaybe<Array<Scalars['String']>>;
  privateIPAddressLT?: InputMaybe<Scalars['String']>;
  privateIPAddressLTE?: InputMaybe<Scalars['String']>;
  privateIPAddressNEQ?: InputMaybe<Scalars['String']>;
  privateIPAddressNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** public_dns_name field predicates */
  publicDNSName?: InputMaybe<Scalars['String']>;
  publicDNSNameContains?: InputMaybe<Scalars['String']>;
  publicDNSNameContainsFold?: InputMaybe<Scalars['String']>;
  publicDNSNameEqualFold?: InputMaybe<Scalars['String']>;
  publicDNSNameGT?: InputMaybe<Scalars['String']>;
  publicDNSNameGTE?: InputMaybe<Scalars['String']>;
  publicDNSNameHasPrefix?: InputMaybe<Scalars['String']>;
  publicDNSNameHasSuffix?: InputMaybe<Scalars['String']>;
  publicDNSNameIn?: InputMaybe<Array<Scalars['String']>>;
  publicDNSNameLT?: InputMaybe<Scalars['String']>;
  publicDNSNameLTE?: InputMaybe<Scalars['String']>;
  publicDNSNameNEQ?: InputMaybe<Scalars['String']>;
  publicDNSNameNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** public_ip_address field predicates */
  publicIPAddress?: InputMaybe<Scalars['String']>;
  publicIPAddressContains?: InputMaybe<Scalars['String']>;
  publicIPAddressContainsFold?: InputMaybe<Scalars['String']>;
  publicIPAddressEqualFold?: InputMaybe<Scalars['String']>;
  publicIPAddressGT?: InputMaybe<Scalars['String']>;
  publicIPAddressGTE?: InputMaybe<Scalars['String']>;
  publicIPAddressHasPrefix?: InputMaybe<Scalars['String']>;
  publicIPAddressHasSuffix?: InputMaybe<Scalars['String']>;
  publicIPAddressIn?: InputMaybe<Array<Scalars['String']>>;
  publicIPAddressLT?: InputMaybe<Scalars['String']>;
  publicIPAddressLTE?: InputMaybe<Scalars['String']>;
  publicIPAddressNEQ?: InputMaybe<Scalars['String']>;
  publicIPAddressNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']>;
  updatedAtGT?: InputMaybe<Scalars['Time']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']>>;
  updatedAtLT?: InputMaybe<Scalars['Time']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']>>;
};

export type Mutation = {
  __typename?: 'Mutation';
  CreateInstance: Instance;
  CreateSlide: Slide;
  CreateSlideWithText: Slide;
  CreateUser: User;
  DeleteInstance: Instance;
  DeleteSlide: Slide;
  DeleteUser: User;
  DeleteUserByEmail: User;
  DeleteUserByUsername: User;
  SignInWithEmail: UserWithAuth;
  SignInWithGithub: UserWithAuth;
  SignInWithUsername: UserWithAuth;
  SignUp: UserWithAuth;
  UpdateInstance: Instance;
  UpdateSlide: Slide;
  UpdateSlideWithText: Slide;
  UpdateUser: User;
};


export type MutationCreateInstanceArgs = {
  input: CreateInstanceInput;
};


export type MutationCreateSlideArgs = {
  input: CreateSlideInput;
};


export type MutationCreateSlideWithTextArgs = {
  input: CreateSlideWithTextInput;
  text: Scalars['String'];
};


export type MutationCreateUserArgs = {
  input: CreateUserInput;
};


export type MutationDeleteInstanceArgs = {
  id: Scalars['ID'];
};


export type MutationDeleteSlideArgs = {
  id: Scalars['ID'];
  user_id: Scalars['ID'];
};


export type MutationDeleteUserArgs = {
  id: Scalars['ID'];
};


export type MutationDeleteUserByEmailArgs = {
  email: Scalars['String'];
};


export type MutationDeleteUserByUsernameArgs = {
  username: Scalars['String'];
};


export type MutationSignInWithEmailArgs = {
  input: SignInWithEmail;
};


export type MutationSignInWithGithubArgs = {
  token: Scalars['String'];
};


export type MutationSignInWithUsernameArgs = {
  input: SignInWithUsername;
};


export type MutationSignUpArgs = {
  input: CreateUserInput;
};


export type MutationUpdateInstanceArgs = {
  input: UpdateInstanceInput;
};


export type MutationUpdateSlideArgs = {
  input: UpdateSlideInput;
};


export type MutationUpdateSlideWithTextArgs = {
  id: Scalars['ID'];
  text: Scalars['String'];
};


export type MutationUpdateUserArgs = {
  input: UpdateUserInput;
};

export type Node = {
  id: Scalars['ID'];
};

export enum OrderDirection {
  Asc = 'ASC',
  Desc = 'DESC'
}

export type PageInfo = {
  __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['Cursor']>;
  hasNextPage: Scalars['Boolean'];
  hasPreviousPage: Scalars['Boolean'];
  startCursor?: Maybe<Scalars['Cursor']>;
};

export type Query = {
  __typename?: 'Query';
  Instance: Instance;
  Instances: InstanceConnection;
  Node?: Maybe<Node>;
  Slide: Slide;
  Slides: SlideConnection;
  User: User;
  UserByAccessToken: UserWithAuth;
  UserByEmail: User;
  UserByUsername: User;
  Users: UserConnection;
  ValidateAccessToken: Scalars['Boolean'];
};


export type QueryInstanceArgs = {
  id: Scalars['ID'];
};


export type QueryInstancesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<InstanceOrder>;
  where?: InputMaybe<InstanceWhereInput>;
};


export type QueryNodeArgs = {
  id: Scalars['ID'];
};


export type QuerySlideArgs = {
  id: Scalars['ID'];
};


export type QuerySlidesArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<SlideOrder>;
  where?: InputMaybe<SlideWhereInput>;
};


export type QueryUserArgs = {
  id: Scalars['ID'];
};


export type QueryUserByAccessTokenArgs = {
  token: Scalars['String'];
};


export type QueryUserByEmailArgs = {
  email: Scalars['String'];
};


export type QueryUserByUsernameArgs = {
  username: Scalars['String'];
};


export type QueryUsersArgs = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  orderBy?: InputMaybe<UserOrder>;
  where?: InputMaybe<UserWhereInput>;
};


export type QueryValidateAccessTokenArgs = {
  token: Scalars['String'];
};

export type SignInWithEmail = {
  email: Scalars['String'];
  password: Scalars['String'];
};

export type SignInWithUsername = {
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Slide = {
  __typename?: 'Slide';
  access_level: AccessLevel;
  created_at: Scalars['Time'];
  deleted: Scalars['Boolean'];
  id: Scalars['ID'];
  name: Scalars['String'];
  path_token?: Maybe<Array<Scalars['String']>>;
  shared_with: Array<Scalars['String']>;
  size?: Maybe<Scalars['Int']>;
  updated_at: Scalars['Time'];
};

export type SlideConnection = {
  __typename?: 'SlideConnection';
  edges: Array<SlideEdge>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type SlideEdge = {
  __typename?: 'SlideEdge';
  cursor: Scalars['Cursor'];
  node: Slide;
};

export type SlideOrder = {
  direction: OrderDirection;
  field?: InputMaybe<SlideOrderField>;
};

export enum SlideOrderField {
  CreatedAt = 'CREATED_AT',
  UpdatedAt = 'UPDATED_AT'
}

/**
 * SlideWhereInput is used for filtering Slide objects.
 * Input was generated by ent.
 */
export type SlideWhereInput = {
  /** access_level field predicates */
  accessLevel?: InputMaybe<AccessLevel>;
  accessLevelIn?: InputMaybe<Array<AccessLevel>>;
  accessLevelNEQ?: InputMaybe<AccessLevel>;
  accessLevelNotIn?: InputMaybe<Array<AccessLevel>>;
  and?: InputMaybe<Array<SlideWhereInput>>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']>;
  createdAtGT?: InputMaybe<Scalars['Time']>;
  createdAtGTE?: InputMaybe<Scalars['Time']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']>>;
  createdAtLT?: InputMaybe<Scalars['Time']>;
  createdAtLTE?: InputMaybe<Scalars['Time']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']>>;
  /** deleted field predicates */
  deleted?: InputMaybe<Scalars['Boolean']>;
  deletedNEQ?: InputMaybe<Scalars['Boolean']>;
  /** instance edge predicates */
  hasInstance?: InputMaybe<Scalars['Boolean']>;
  hasInstanceWith?: InputMaybe<Array<InstanceWhereInput>>;
  /** user edge predicates */
  hasUser?: InputMaybe<Scalars['Boolean']>;
  hasUserWith?: InputMaybe<Array<UserWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']>;
  idGT?: InputMaybe<Scalars['ID']>;
  idGTE?: InputMaybe<Scalars['ID']>;
  idIn?: InputMaybe<Array<Scalars['ID']>>;
  idLT?: InputMaybe<Scalars['ID']>;
  idLTE?: InputMaybe<Scalars['ID']>;
  idNEQ?: InputMaybe<Scalars['ID']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']>;
  nameContains?: InputMaybe<Scalars['String']>;
  nameContainsFold?: InputMaybe<Scalars['String']>;
  nameEqualFold?: InputMaybe<Scalars['String']>;
  nameGT?: InputMaybe<Scalars['String']>;
  nameGTE?: InputMaybe<Scalars['String']>;
  nameHasPrefix?: InputMaybe<Scalars['String']>;
  nameHasSuffix?: InputMaybe<Scalars['String']>;
  nameIn?: InputMaybe<Array<Scalars['String']>>;
  nameLT?: InputMaybe<Scalars['String']>;
  nameLTE?: InputMaybe<Scalars['String']>;
  nameNEQ?: InputMaybe<Scalars['String']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']>>;
  not?: InputMaybe<SlideWhereInput>;
  or?: InputMaybe<Array<SlideWhereInput>>;
  /** size field predicates */
  size?: InputMaybe<Scalars['Int']>;
  sizeGT?: InputMaybe<Scalars['Int']>;
  sizeGTE?: InputMaybe<Scalars['Int']>;
  sizeIn?: InputMaybe<Array<Scalars['Int']>>;
  sizeIsNil?: InputMaybe<Scalars['Boolean']>;
  sizeLT?: InputMaybe<Scalars['Int']>;
  sizeLTE?: InputMaybe<Scalars['Int']>;
  sizeNEQ?: InputMaybe<Scalars['Int']>;
  sizeNotIn?: InputMaybe<Array<Scalars['Int']>>;
  sizeNotNil?: InputMaybe<Scalars['Boolean']>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']>;
  updatedAtGT?: InputMaybe<Scalars['Time']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']>>;
  updatedAtLT?: InputMaybe<Scalars['Time']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']>>;
};

export type UpdateInstanceInput = {
  architecture?: InputMaybe<Scalars['String']>;
  availability_zone?: InputMaybe<Scalars['String']>;
  clear_slide?: InputMaybe<Scalars['Boolean']>;
  clear_user?: InputMaybe<Scalars['Boolean']>;
  id: Scalars['ID'];
  image_id?: InputMaybe<Scalars['String']>;
  instance_type?: InputMaybe<Scalars['String']>;
  private_dns_name?: InputMaybe<Scalars['String']>;
  private_ip_address?: InputMaybe<Scalars['String']>;
  public_dns_name?: InputMaybe<Scalars['String']>;
  public_ip_address?: InputMaybe<Scalars['String']>;
  slide_id?: InputMaybe<Scalars['ID']>;
  user_id?: InputMaybe<Scalars['ID']>;
};

export type UpdateSlideInput = {
  access_level?: InputMaybe<AccessLevel>;
  clear_instance?: InputMaybe<Scalars['Boolean']>;
  clear_path_token?: InputMaybe<Scalars['Boolean']>;
  clear_user?: InputMaybe<Scalars['Boolean']>;
  deleted?: InputMaybe<Scalars['Boolean']>;
  id: Scalars['ID'];
  instance_id?: InputMaybe<Scalars['ID']>;
  name?: InputMaybe<Scalars['String']>;
  path_token?: InputMaybe<Array<Scalars['String']>>;
  shared_with?: InputMaybe<Array<Scalars['String']>>;
  size?: InputMaybe<Scalars['Int']>;
  user_id?: InputMaybe<Scalars['ID']>;
};

export type UpdateUserInput = {
  avatar_url?: InputMaybe<Scalars['String']>;
  bio?: InputMaybe<Scalars['String']>;
  clear_avatar_url?: InputMaybe<Scalars['Boolean']>;
  clear_bio?: InputMaybe<Scalars['Boolean']>;
  full_name?: InputMaybe<Scalars['String']>;
  id: Scalars['ID'];
  username?: InputMaybe<Scalars['String']>;
};

export type User = {
  __typename?: 'User';
  avatar_url?: Maybe<Scalars['String']>;
  bio?: Maybe<Scalars['String']>;
  created_at: Scalars['Time'];
  email: Scalars['String'];
  full_name?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  updated_at: Scalars['Time'];
  username: Scalars['String'];
};

export type UserConnection = {
  __typename?: 'UserConnection';
  edges?: Maybe<Array<Maybe<UserEdge>>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type UserEdge = {
  __typename?: 'UserEdge';
  cursor: Scalars['Cursor'];
  node?: Maybe<User>;
};

export type UserOrder = {
  direction: OrderDirection;
  field?: InputMaybe<UserOrderField>;
};

export enum UserOrderField {
  CreatedAt = 'CREATED_AT',
  UpdatedAt = 'UPDATED_AT'
}

/**
 * UserWhereInput is used for filtering User objects.
 * Input was generated by ent.
 */
export type UserWhereInput = {
  and?: InputMaybe<Array<UserWhereInput>>;
  /** avatar_url field predicates */
  avatarURL?: InputMaybe<Scalars['String']>;
  avatarURLContains?: InputMaybe<Scalars['String']>;
  avatarURLContainsFold?: InputMaybe<Scalars['String']>;
  avatarURLEqualFold?: InputMaybe<Scalars['String']>;
  avatarURLGT?: InputMaybe<Scalars['String']>;
  avatarURLGTE?: InputMaybe<Scalars['String']>;
  avatarURLHasPrefix?: InputMaybe<Scalars['String']>;
  avatarURLHasSuffix?: InputMaybe<Scalars['String']>;
  avatarURLIn?: InputMaybe<Array<Scalars['String']>>;
  avatarURLIsNil?: InputMaybe<Scalars['Boolean']>;
  avatarURLLT?: InputMaybe<Scalars['String']>;
  avatarURLLTE?: InputMaybe<Scalars['String']>;
  avatarURLNEQ?: InputMaybe<Scalars['String']>;
  avatarURLNotIn?: InputMaybe<Array<Scalars['String']>>;
  avatarURLNotNil?: InputMaybe<Scalars['Boolean']>;
  /** bio field predicates */
  bio?: InputMaybe<Scalars['String']>;
  bioContains?: InputMaybe<Scalars['String']>;
  bioContainsFold?: InputMaybe<Scalars['String']>;
  bioEqualFold?: InputMaybe<Scalars['String']>;
  bioGT?: InputMaybe<Scalars['String']>;
  bioGTE?: InputMaybe<Scalars['String']>;
  bioHasPrefix?: InputMaybe<Scalars['String']>;
  bioHasSuffix?: InputMaybe<Scalars['String']>;
  bioIn?: InputMaybe<Array<Scalars['String']>>;
  bioIsNil?: InputMaybe<Scalars['Boolean']>;
  bioLT?: InputMaybe<Scalars['String']>;
  bioLTE?: InputMaybe<Scalars['String']>;
  bioNEQ?: InputMaybe<Scalars['String']>;
  bioNotIn?: InputMaybe<Array<Scalars['String']>>;
  bioNotNil?: InputMaybe<Scalars['Boolean']>;
  /** created_at field predicates */
  createdAt?: InputMaybe<Scalars['Time']>;
  createdAtGT?: InputMaybe<Scalars['Time']>;
  createdAtGTE?: InputMaybe<Scalars['Time']>;
  createdAtIn?: InputMaybe<Array<Scalars['Time']>>;
  createdAtLT?: InputMaybe<Scalars['Time']>;
  createdAtLTE?: InputMaybe<Scalars['Time']>;
  createdAtNEQ?: InputMaybe<Scalars['Time']>;
  createdAtNotIn?: InputMaybe<Array<Scalars['Time']>>;
  /** email field predicates */
  email?: InputMaybe<Scalars['String']>;
  emailContains?: InputMaybe<Scalars['String']>;
  emailContainsFold?: InputMaybe<Scalars['String']>;
  emailEqualFold?: InputMaybe<Scalars['String']>;
  emailGT?: InputMaybe<Scalars['String']>;
  emailGTE?: InputMaybe<Scalars['String']>;
  emailHasPrefix?: InputMaybe<Scalars['String']>;
  emailHasSuffix?: InputMaybe<Scalars['String']>;
  emailIn?: InputMaybe<Array<Scalars['String']>>;
  emailLT?: InputMaybe<Scalars['String']>;
  emailLTE?: InputMaybe<Scalars['String']>;
  emailNEQ?: InputMaybe<Scalars['String']>;
  emailNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** full_name field predicates */
  fullName?: InputMaybe<Scalars['String']>;
  fullNameContains?: InputMaybe<Scalars['String']>;
  fullNameContainsFold?: InputMaybe<Scalars['String']>;
  fullNameEqualFold?: InputMaybe<Scalars['String']>;
  fullNameGT?: InputMaybe<Scalars['String']>;
  fullNameGTE?: InputMaybe<Scalars['String']>;
  fullNameHasPrefix?: InputMaybe<Scalars['String']>;
  fullNameHasSuffix?: InputMaybe<Scalars['String']>;
  fullNameIn?: InputMaybe<Array<Scalars['String']>>;
  fullNameIsNil?: InputMaybe<Scalars['Boolean']>;
  fullNameLT?: InputMaybe<Scalars['String']>;
  fullNameLTE?: InputMaybe<Scalars['String']>;
  fullNameNEQ?: InputMaybe<Scalars['String']>;
  fullNameNotIn?: InputMaybe<Array<Scalars['String']>>;
  fullNameNotNil?: InputMaybe<Scalars['Boolean']>;
  /** instances edge predicates */
  hasInstances?: InputMaybe<Scalars['Boolean']>;
  hasInstancesWith?: InputMaybe<Array<InstanceWhereInput>>;
  /** slides edge predicates */
  hasSlides?: InputMaybe<Scalars['Boolean']>;
  hasSlidesWith?: InputMaybe<Array<SlideWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']>;
  idGT?: InputMaybe<Scalars['ID']>;
  idGTE?: InputMaybe<Scalars['ID']>;
  idIn?: InputMaybe<Array<Scalars['ID']>>;
  idLT?: InputMaybe<Scalars['ID']>;
  idLTE?: InputMaybe<Scalars['ID']>;
  idNEQ?: InputMaybe<Scalars['ID']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']>>;
  not?: InputMaybe<UserWhereInput>;
  or?: InputMaybe<Array<UserWhereInput>>;
  /** password_hash field predicates */
  passwordHash?: InputMaybe<Scalars['String']>;
  passwordHashContains?: InputMaybe<Scalars['String']>;
  passwordHashContainsFold?: InputMaybe<Scalars['String']>;
  passwordHashEqualFold?: InputMaybe<Scalars['String']>;
  passwordHashGT?: InputMaybe<Scalars['String']>;
  passwordHashGTE?: InputMaybe<Scalars['String']>;
  passwordHashHasPrefix?: InputMaybe<Scalars['String']>;
  passwordHashHasSuffix?: InputMaybe<Scalars['String']>;
  passwordHashIn?: InputMaybe<Array<Scalars['String']>>;
  passwordHashLT?: InputMaybe<Scalars['String']>;
  passwordHashLTE?: InputMaybe<Scalars['String']>;
  passwordHashNEQ?: InputMaybe<Scalars['String']>;
  passwordHashNotIn?: InputMaybe<Array<Scalars['String']>>;
  /** updated_at field predicates */
  updatedAt?: InputMaybe<Scalars['Time']>;
  updatedAtGT?: InputMaybe<Scalars['Time']>;
  updatedAtGTE?: InputMaybe<Scalars['Time']>;
  updatedAtIn?: InputMaybe<Array<Scalars['Time']>>;
  updatedAtLT?: InputMaybe<Scalars['Time']>;
  updatedAtLTE?: InputMaybe<Scalars['Time']>;
  updatedAtNEQ?: InputMaybe<Scalars['Time']>;
  updatedAtNotIn?: InputMaybe<Array<Scalars['Time']>>;
  /** username field predicates */
  username?: InputMaybe<Scalars['String']>;
  usernameContains?: InputMaybe<Scalars['String']>;
  usernameContainsFold?: InputMaybe<Scalars['String']>;
  usernameEqualFold?: InputMaybe<Scalars['String']>;
  usernameGT?: InputMaybe<Scalars['String']>;
  usernameGTE?: InputMaybe<Scalars['String']>;
  usernameHasPrefix?: InputMaybe<Scalars['String']>;
  usernameHasSuffix?: InputMaybe<Scalars['String']>;
  usernameIn?: InputMaybe<Array<Scalars['String']>>;
  usernameLT?: InputMaybe<Scalars['String']>;
  usernameLTE?: InputMaybe<Scalars['String']>;
  usernameNEQ?: InputMaybe<Scalars['String']>;
  usernameNotIn?: InputMaybe<Array<Scalars['String']>>;
};

export type UserWithAuth = {
  __typename?: 'UserWithAuth';
  access_token: Scalars['String'];
  expired_at: Scalars['Time'];
  user: User;
};

export type UserByAccessTokenQueryVariables = Exact<{
  token: Scalars['String'];
}>;


export type UserByAccessTokenQuery = { __typename?: 'Query', UserByAccessToken: { __typename?: 'UserWithAuth', access_token: string, expired_at: any, user: { __typename?: 'User', id: string, username: string, email: string, full_name?: string | null, avatar_url?: string | null, bio?: string | null, created_at: any, updated_at: any } } };

export type SignUpMutationVariables = Exact<{
  input: CreateUserInput;
}>;


export type SignUpMutation = { __typename?: 'Mutation', SignUp: { __typename?: 'UserWithAuth', access_token: string, expired_at: any, user: { __typename?: 'User', id: string, username: string, email: string, full_name?: string | null, avatar_url?: string | null, bio?: string | null, created_at: any, updated_at: any } } };

export type SignInWithUsernameMutationVariables = Exact<{
  input: SignInWithUsername;
}>;


export type SignInWithUsernameMutation = { __typename?: 'Mutation', SignInWithUsername: { __typename?: 'UserWithAuth', access_token: string, expired_at: any, user: { __typename?: 'User', id: string, username: string, email: string, full_name?: string | null, avatar_url?: string | null, bio?: string | null, created_at: any, updated_at: any } } };

export type SignInWithEmailMutationVariables = Exact<{
  input: SignInWithEmail;
}>;


export type SignInWithEmailMutation = { __typename?: 'Mutation', SignInWithEmail: { __typename?: 'UserWithAuth', access_token: string, expired_at: any, user: { __typename?: 'User', id: string, username: string, email: string, full_name?: string | null, avatar_url?: string | null, bio?: string | null, created_at: any, updated_at: any } } };

export type UserFieldsFragment = { __typename?: 'User', id: string, username: string, email: string, full_name?: string | null, avatar_url?: string | null, bio?: string | null, created_at: any, updated_at: any };

export type UserWithAuthFieldsFragment = { __typename?: 'UserWithAuth', access_token: string, expired_at: any, user: { __typename?: 'User', id: string, username: string, email: string, full_name?: string | null, avatar_url?: string | null, bio?: string | null, created_at: any, updated_at: any } };

export type GetSlideQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetSlideQuery = { __typename?: 'Query', Slide: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } };

export type ListSlideQueryVariables = Exact<{
  after?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  before?: InputMaybe<Scalars['Cursor']>;
  last?: InputMaybe<Scalars['Int']>;
  where?: InputMaybe<SlideWhereInput>;
  orderBy?: InputMaybe<SlideOrder>;
}>;


export type ListSlideQuery = { __typename?: 'Query', Slides: { __typename?: 'SlideConnection', totalCount: number, pageInfo: { __typename?: 'PageInfo', hasNextPage: boolean, hasPreviousPage: boolean, startCursor?: any | null, endCursor?: any | null }, edges: Array<{ __typename?: 'SlideEdge', cursor: any, node: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } }> } };

export type CreateSlideMutationVariables = Exact<{
  input: CreateSlideInput;
}>;


export type CreateSlideMutation = { __typename?: 'Mutation', CreateSlide: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } };

export type CreateSlideWithTextMutationVariables = Exact<{
  createSlideWithTextInput: CreateSlideWithTextInput;
  text: Scalars['String'];
}>;


export type CreateSlideWithTextMutation = { __typename?: 'Mutation', CreateSlideWithText: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } };

export type UpdateSlideWithTextMutationVariables = Exact<{
  id: Scalars['ID'];
  text: Scalars['String'];
}>;


export type UpdateSlideWithTextMutation = { __typename?: 'Mutation', UpdateSlideWithText: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } };

export type UpdateSlideMutationVariables = Exact<{
  input: UpdateSlideInput;
}>;


export type UpdateSlideMutation = { __typename?: 'Mutation', UpdateSlide: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } };

export type DeleteSlideMutationVariables = Exact<{
  id: Scalars['ID'];
  user_id: Scalars['ID'];
}>;


export type DeleteSlideMutation = { __typename?: 'Mutation', DeleteSlide: { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any } };

export type SlideFieldsFragment = { __typename?: 'Slide', id: string, name: string, path_token?: Array<string> | null, size?: number | null, access_level: AccessLevel, shared_with: Array<string>, deleted: boolean, created_at: any, updated_at: any };

export const UserFieldsFragmentDoc = gql`
    fragment UserFields on User {
  id
  username
  email
  full_name
  avatar_url
  bio
  created_at
  updated_at
}
    `;
export const UserWithAuthFieldsFragmentDoc = gql`
    fragment UserWithAuthFields on UserWithAuth {
  access_token
  expired_at
  user {
    ...UserFields
  }
}
    ${UserFieldsFragmentDoc}`;
export const SlideFieldsFragmentDoc = gql`
    fragment SlideFields on Slide {
  id
  name
  path_token
  size
  access_level
  shared_with
  deleted
  created_at
  updated_at
}
    `;
export const UserByAccessTokenDocument = gql`
    query UserByAccessToken($token: String!) {
  UserByAccessToken(token: $token) {
    ...UserWithAuthFields
  }
}
    ${UserWithAuthFieldsFragmentDoc}`;

export function useUserByAccessTokenQuery(options: Omit<Urql.UseQueryArgs<UserByAccessTokenQueryVariables>, 'query'>) {
  return Urql.useQuery<UserByAccessTokenQuery>({ query: UserByAccessTokenDocument, ...options });
};
export const SignUpDocument = gql`
    mutation SignUp($input: CreateUserInput!) {
  SignUp(input: $input) {
    ...UserWithAuthFields
  }
}
    ${UserWithAuthFieldsFragmentDoc}`;

export function useSignUpMutation() {
  return Urql.useMutation<SignUpMutation, SignUpMutationVariables>(SignUpDocument);
};
export const SignInWithUsernameDocument = gql`
    mutation SignInWithUsername($input: SignInWithUsername!) {
  SignInWithUsername(input: $input) {
    ...UserWithAuthFields
  }
}
    ${UserWithAuthFieldsFragmentDoc}`;

export function useSignInWithUsernameMutation() {
  return Urql.useMutation<SignInWithUsernameMutation, SignInWithUsernameMutationVariables>(SignInWithUsernameDocument);
};
export const SignInWithEmailDocument = gql`
    mutation SignInWithEmail($input: SignInWithEmail!) {
  SignInWithEmail(input: $input) {
    ...UserWithAuthFields
  }
}
    ${UserWithAuthFieldsFragmentDoc}`;

export function useSignInWithEmailMutation() {
  return Urql.useMutation<SignInWithEmailMutation, SignInWithEmailMutationVariables>(SignInWithEmailDocument);
};
export const GetSlideDocument = gql`
    query GetSlide($id: ID!) {
  Slide(id: $id) {
    ...SlideFields
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useGetSlideQuery(options: Omit<Urql.UseQueryArgs<GetSlideQueryVariables>, 'query'>) {
  return Urql.useQuery<GetSlideQuery>({ query: GetSlideDocument, ...options });
};
export const ListSlideDocument = gql`
    query ListSlide($after: Cursor, $first: Int, $before: Cursor, $last: Int, $where: SlideWhereInput, $orderBy: SlideOrder) {
  Slides(
    after: $after
    first: $first
    before: $before
    last: $last
    where: $where
    orderBy: $orderBy
  ) {
    totalCount
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    edges {
      node {
        ...SlideFields
      }
      cursor
    }
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useListSlideQuery(options?: Omit<Urql.UseQueryArgs<ListSlideQueryVariables>, 'query'>) {
  return Urql.useQuery<ListSlideQuery>({ query: ListSlideDocument, ...options });
};
export const CreateSlideDocument = gql`
    mutation CreateSlide($input: CreateSlideInput!) {
  CreateSlide(input: $input) {
    ...SlideFields
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useCreateSlideMutation() {
  return Urql.useMutation<CreateSlideMutation, CreateSlideMutationVariables>(CreateSlideDocument);
};
export const CreateSlideWithTextDocument = gql`
    mutation CreateSlideWithText($createSlideWithTextInput: CreateSlideWithTextInput!, $text: String!) {
  CreateSlideWithText(input: $createSlideWithTextInput, text: $text) {
    ...SlideFields
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useCreateSlideWithTextMutation() {
  return Urql.useMutation<CreateSlideWithTextMutation, CreateSlideWithTextMutationVariables>(CreateSlideWithTextDocument);
};
export const UpdateSlideWithTextDocument = gql`
    mutation UpdateSlideWithText($id: ID!, $text: String!) {
  UpdateSlideWithText(id: $id, text: $text) {
    ...SlideFields
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useUpdateSlideWithTextMutation() {
  return Urql.useMutation<UpdateSlideWithTextMutation, UpdateSlideWithTextMutationVariables>(UpdateSlideWithTextDocument);
};
export const UpdateSlideDocument = gql`
    mutation UpdateSlide($input: UpdateSlideInput!) {
  UpdateSlide(input: $input) {
    ...SlideFields
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useUpdateSlideMutation() {
  return Urql.useMutation<UpdateSlideMutation, UpdateSlideMutationVariables>(UpdateSlideDocument);
};
export const DeleteSlideDocument = gql`
    mutation DeleteSlide($id: ID!, $user_id: ID!) {
  DeleteSlide(id: $id, user_id: $user_id) {
    ...SlideFields
  }
}
    ${SlideFieldsFragmentDoc}`;

export function useDeleteSlideMutation() {
  return Urql.useMutation<DeleteSlideMutation, DeleteSlideMutationVariables>(DeleteSlideDocument);
};