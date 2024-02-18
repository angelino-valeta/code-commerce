import { Injectable, UnauthorizedException } from '@nestjs/common'
import { JwtService } from '@nestjs/jwt';

const users = [
  {
    id: 1,
    username: 'angelino',
    password: 'angelino',
  },
  {
    id: 2,
    username: 'valeta',
    password: 'valeta',
  },
];

@Injectable()
export class AuthService {
  constructor(private jwtService: JwtService) { }

  login(username: string, password: string) {
    const user = users.find(
      (user) => user.username === username && user.password === password
    );

    if (!user) {
      throw new UnauthorizedException();
    }

    const payload = { sub: user.id, username: user.username };

    return {
      access_token: this.jwtService.sign(payload),
    };
  }
}
