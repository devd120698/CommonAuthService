# CommonAuthService
This repo will provide APIs for Authorization and Authentication for users, Later on I might add user roles as well

### Overview

This will be a microservice, which will have set of APIs that will one stop place that enables User creation, SignIn, Sign Up of users, role based handling.


## DB Schema

#### 1. Users

```
CREATE TABLE `Users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phoneNo` varchar(25) DEFAULT NULL,
  `addedOn` datetime DEFAULT NULL,
  `updatedOn` datetime DEFAULT NULL,
  `encPassword` text,
  `address` varchar(100) DEFAULT NULL,
  `role` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'USER',
  `isActive` tinyint(1) DEFAULT NULL,
  `profileImage` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

```

#### 2. Notes

```
id
noteId
title
note -> index(research on what is the best way to store and retrieve strings) 
userId -> foreign key to Users table
attachments
```



## API and Feature Specs

#### Potential APIs

1. Create User
    1. Define and enforce password complexity rules
2. Deactivate User -> Only accessible to admin
3. Get User by emailId -> Only accessible to admin
4. Sign In API -> Only accessible to admin and user(Send back JWT token) -> rest all API calls should have that token, and use that token to verify the user
5. Verify token API
6. Sign out API -> Only accessible to admin and user(send)
7. Update Password -> Only accessible to admin and user
    1. Admin can update their own password
    2. User can update their own password
8. Update User details -> Profile Image, name, address, phoneNo
9. Add notes (first verify the token and add, return error if not valid)

#### Potential Features
1. Add multi factor auth
    1. TOTP
    2. OAuth
    3. Google Authenticator
2. Audit Logs
    1. Store User table updates in mongo or someplace




# Features to look into


Creating a user authentication and authorization platform in Go is a great project idea that can be both technically challenging and useful to others. Here are some features you can consider implementing to make your project more comprehensive and technically complex:

1. Role-based Access Control (RBAC):

    - Implement a robust RBAC system that allows you to define roles and permissions for users and resources.
    - Enable role-based authorization for various API endpoints and actions.
2. Multi-factor Authentication (MFA):

    - Implement support for MFA to enhance security.
    - Include options for using TOTP (Time-based One-Time Password) and SMS-based MFA.
3. OAuth and OpenID Connect:

    - Allow users to authenticate using popular third-party identity providers such as Google, Facebook, or GitHub.
    - Implement OAuth 2.0 and OpenID Connect for secure authorization and user profile retrieval.
4. JWT (JSON Web Tokens):

    - Utilize JWTs for token-based authentication and authorization.
    - Implement token issuance, validation, and revocation mechanisms.
5. User Profile Management:

    - Enhance user profile management by allowing users to update their profiles, reset passwords, or manage email preferences.
6. Audit Logs:

    - Implement a comprehensive auditing system to log user actions, authentication events, and authorization decisions for compliance and security.
7. Password Policies:

    - Define and enforce password complexity rules.
    - Implement password hashing and salting for security.
8. Rate Limiting and IP Whitelisting/Blacklisting:

    - Protect your authentication and authorization endpoints with rate limiting to prevent abuse.
    - Allow administrators to whitelist or blacklist IP addresses.
9. Single Sign-On (SSO):

    - Support SSO for a seamless user experience across multiple applications.
    - Implement popular SSO protocols like SAML or OpenID Connect.
10. Extensibility:

    - Design your platform to be highly extensible, allowing others to integrate custom authentication methods, connectors to various databases, and custom actions.
11. Account Linking:

    - Allow users to link multiple authentication methods (e.g., email/password, social login) to a single account.
12. Multi-Tenant Support:

    - Enable multi-tenancy, so organizations can manage their users and permissions within a shared system.
13. Internationalization and Localization:

    - Make your platform accessible to a global audience by supporting multiple languages and cultural preferences.
14. Secure Token Storage:

    - Ensure secure storage of tokens, API keys, and other sensitive data.
15. API Documentation and Developer Resources:

    - Provide comprehensive API documentation, code examples, and developer resources to encourage adoption and usage.
16. Performance Optimization and Caching:

    - Optimize the performance of your authentication and authorization system through efficient database queries, caching, and load balancing.
17. Containerization and Deployment:

    - Dockerize your application for easy deployment in various cloud environments, making it more accessible to developers.
18. Open Source and Community Support:

    - Consider making your project open source to encourage contributions and community support.