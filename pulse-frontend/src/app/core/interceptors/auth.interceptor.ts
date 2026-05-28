import { HttpInterceptorFn } from '@angular/common/http';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  // 1. Get the token from local storage
  const token = localStorage.getItem('authToken');

  // 2. If we have a token, clone the request and add the Authorization header
  if (token) {
    const clonedRequest = req.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`
      }
    });
    // 3. Send the cloned request onward
    return next(clonedRequest);
  }

  // 4. If no token, just send the original request
  return next(req);
};
