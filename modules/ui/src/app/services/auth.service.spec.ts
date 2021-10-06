import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { environment } from '@environments/environment';
import { AuthService } from '@services/auth.service';
import { MockRouter } from '@testing/mock-router';

describe('AuthService', () => {
    let service: AuthService;
    let httpMock: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [
                HttpClientTestingModule,
            ],
            providers: [
                { provide: Router, useClass: MockRouter },
            ],
        });

        httpMock = TestBed.inject(HttpTestingController);
        service = TestBed.inject(AuthService);
    });

    it('should be created', () => {
        expect(service).toBeTruthy();
    });

    it('should set isLoggedIn to true on login', () => {
        const token = 'some_access_token';
        service.login('username', 'password').toPromise();

        const req = httpMock.expectOne(
            (request) =>
                request.method === 'POST' &&
                request.url === `${environment.apiUrl}/login`);
        req.flush({access_token: token});
        httpMock.verify();
        expect(service.isLoggedIn).toEqual(true);
    });

    it('should not allow activation if logged in', () => {
        service.isLoggedIn = true;
        expect(service.canActivate()).toEqual(false);
    });

    it('should allow activation if not logged in', () => {
        service.isLoggedIn = false;
        expect(service.canActivate()).toEqual(true);
    });

    it('should logout', () => {
        service.isLoggedIn = true;
        service.logout();
        expect(service.isLoggedIn).toEqual(false);
    });

    const checkGetToken = (token: string|null, auth: boolean) => {
        service.isLoggedIn = auth;
        spyOn(localStorage, 'getItem').and.returnValue(token);
        expect(service.getToken()).toEqual(token);
        expect(service.isLoggedIn).toEqual(!auth);
    };
    it('should logout if no token in storage', () => {
        checkGetToken(null, true);
    });

    it('should login if a token in storage', () => {
        checkGetToken('some_token', false);
    });
});
