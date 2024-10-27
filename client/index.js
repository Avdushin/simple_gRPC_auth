const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

//? Загрузка .proto файла
const packageDefinition = protoLoader.loadSync('proto/user.proto', {});
const userProto = grpc.loadPackageDefinition(packageDefinition).pb;

const client = new userProto.UserService('localhost:50051', grpc.credentials.createInsecure());

//@ SignUp
const registerUser = () => {
  const userData = {
    username: 'testuser',
    email: 'test@example.com',
    password: 'password123',
  };

  client.RegisterUser(userData, (error, response) => {
    if (error) {
      console.error('Error:', error);
    } else {
      console.log('Register response:', response);
    }
  });
};

//@ Login
const loginUser = () => {
  const loginData = {
    email: 'test@example.com',
    password: 'password123',
  };

  client.LoginUser(loginData, (error, response) => {
    if (error) {
      console.error('Error:', error);
    } else {
      console.log('Login response:', response);
    }
  });
};

registerUser();
setTimeout(() => {
  loginUser();
}, 2000);
